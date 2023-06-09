function getLocalFilePath(path) {
	// 检查路径是否以指定关键词开头
	if (path.indexOf('_www') === 0 || path.indexOf('_doc') === 0 || path.indexOf('_documents') === 0 || path.indexOf(
			'_downloads') === 0) {
		return path; // 如果路径以指定关键词开头，则直接返回该路径
	}
	// 检查路径是否以 'file://' 开头
	if (path.indexOf('file://') === 0) {
		return path; // 如果路径以 'file://' 开头，则直接返回该路径
	}
	// 检查路径是否以 '/storage/emulated/0/' 开头
	if (path.indexOf('/storage/emulated/0/') === 0) {
		return path; // 如果路径以 '/storage/emulated/0/' 开头，则直接返回该路径
	}
	// 检查路径是否以 '/' 开头
	if (path.indexOf('/') === 0) {
		// 将绝对文件系统路径转换为本地路径
		var localFilePath = plus.io.convertAbsoluteFileSystem(path);
		if (localFilePath !== path) {
			return localFilePath; // 如果转换后的路径与原始路径不同，则返回转换后的路径
		} else {
			path = path.substr(1); // 如果转换后的路径与原始路径相同，则去掉路径开头的 '/'
		}
	}
	// 默认返回路径以 '_www/' 开头
	return '_www/' + path;
}


// 将 data URL 转换为 base64 字符串
function dataUrlToBase64(str) {
	var array = str.split(','); // 将字符串按逗号分隔为数组
	return array[array.length - 1]; // 返回数组最后一个元素，即 base64 字符串部分
}

var index = 0; // 初始化索引变量
// 获取一个新的文件 ID
function getNewFileId() {
	return Date.now() + String(index++); // 使用当前时间戳和索引构建一个唯一的文件 ID
}

// 比较两个版本号大小
function biggerThan(v1, v2) {
	var v1Array = v1.split('.'); // 将版本号 v1 按点号分隔为数组
	var v2Array = v2.split('.'); // 将版本号 v2 按点号分隔为数组
	var update = false; // 初始化更新标志为 false
	for (var index = 0; index < v2Array.length; index++) {
		var diff = v1Array[index] - v2Array[index]; // 比较相同位置的版本号部分的差异
		if (diff !== 0) {
			update = diff > 0; // 如果差异不为零，根据差值的正负确定是否需要更新
			break;
		}
	}
	return update; // 返回更新标志
}


export function pathToBase64(path) {
	return new Promise(function(resolve, reject) {
		if (typeof window === 'object' && 'document' in window) {
			// 如果在浏览器环境中
			if (typeof FileReader === 'function') {
				// 检查浏览器是否支持FileReader
				var xhr = new XMLHttpRequest();
				xhr.open('GET', path, true);
				xhr.responseType = 'blob';
				xhr.onload = function() {
					if (this.status === 200) {
						let fileReader = new FileReader();
						fileReader.onload = function(e) {
							resolve(e.target.result);
						};
						fileReader.onerror = reject;
						fileReader.readAsDataURL(this.response);
					}
				};
				xhr.onerror = reject;
				xhr.send();
				return;
			}
			// 如果浏览器不支持FileReader，则使用canvas进行处理
			var canvas = document.createElement('canvas');
			var c2x = canvas.getContext('2d');
			var img = new Image();
			img.onload = function() {
				canvas.width = img.width;
				canvas.height = img.height;
				c2x.drawImage(img, 0, 0);
				resolve(canvas.toDataURL());
				canvas.height = canvas.width = 0;
			};
			img.onerror = reject;
			img.src = path;
			return;
		}
		// 如果是在plus环境中
		if (typeof plus === 'object') {
			plus.io.resolveLocalFileSystemURL(getLocalFilePath(path), function(entry) {
				entry.file(function(file) {
					var fileReader = new plus.io.FileReader();
					fileReader.onload = function(data) {
						resolve(data.target.result);
					};
					fileReader.onerror = function(error) {
						reject(error);
					};
					fileReader.readAsDataURL(file);
				}, function(error) {
					reject(error);
				});
			}, function(error) {
				reject(error);
			});
			return;
		}
		// 如果是在微信小程序环境中
		if (typeof wx === 'object' && wx.canIUse('getFileSystemManager')) {
			wx.getFileSystemManager().readFile({
				filePath: path,
				encoding: 'base64',
				success: function(res) {
					resolve('data:image/png;base64,' + res.data);
				},
				fail: function(error) {
					reject(error);
				}
			});
			return;
		}
		// 如果不支持以上环境，则抛出错误
		reject(new Error('not support'));
	});
}

export function base64ToPath(base64) {
	return new Promise(function(resolve, reject) {
		if (typeof window === 'object' && 'document' in window) {
			// 如果在浏览器环境中
			base64 = base64.split(',');
			var type = base64[0].match(/:(.*?);/)[1];
			var str = atob(base64[1]);
			var n = str.length;
			var array = new Uint8Array(n);
			while (n--) {
				array[n] = str.charCodeAt(n);
			}
			return resolve((window.URL || window.webkitURL).createObjectURL(new Blob([array], {
				type: type
			})));
		}
		var extName = base64.split(',')[0].match(/data\:\S+\/(\S+);/);
		if (extName) {
			extName = extName[1];
		} else {
			reject(new Error('base64 error'));
		}
		var fileName = getNewFileId() + '.' + extName;
		if (typeof plus === 'object') {
			// 如果是在plus环境中
			var basePath = '_doc';
			var dirPath = 'uniapp_temp';
			var filePath = basePath + '/' + dirPath + '/' + fileName;
			if (!biggerThan(plus.os.name === 'Android' ? '1.9.9.80627' : '1.9.9.80472', plus.runtime
					.innerVersion)) {
				// 如果plus版本较低，使用旧版plus.io接口
				plus.io.resolveLocalFileSystemURL(basePath, function(entry) {
					entry.getDirectory(dirPath, {
						create: true,
						exclusive: false,
					}, function(entry) {
						entry.getFile(fileName, {
							create: true,
							exclusive: false,
						}, function(entry) {
							entry.createWriter(function(writer) {
								writer.onwrite = function() {
									resolve(filePath);
								};
								writer.onerror = reject;
								writer.seek(0);
								writer.writeAsBinary(dataUrlToBase64(base64));
							}, reject);
						}, reject);
					}, reject);
				}, reject);
				return;
			}
			// 如果plus版本较高，使用新版plus.nativeObj.Bitmap接口
			var bitmap = new plus.nativeObj.Bitmap(fileName);
			bitmap.loadBase64Data(base64, function() {
				bitmap.save(filePath, {}, function() {
					bitmap.clear();
					resolve(filePath);
				}, function(error) {
					bitmap.clear();
					reject(error);
				});
			}, function(error) {
				bitmap.clear();
				reject(error);
			});
			return;
		}
		// 如果是在微信小程序环境中
		if (typeof wx === 'object' && wx.canIUse('getFileSystemManager')) {
			var filePath = wx.env.USER_DATA_PATH + '/' + fileName;
			wx.getFileSystemManager().writeFile({
				filePath: filePath,
				data: dataUrlToBase64(base64),
				encoding: 'base64',
				success: function() {
					resolve(filePath);
				},
				fail: function(error) {
					reject(error);
				}
			});
			return;
		}
		// 如果不支持以上环境，则抛出错误
		reject(new Error('not support'));
	});
}
