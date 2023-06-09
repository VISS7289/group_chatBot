export default {
	//转换时间
	saveBase64Img: function(base64) {
		const bitmap = new plus.nativeObj.Bitmap('test'); // 创建一个Bitmap对象
		bitmap.loadBase64Data(
			base64, // 要加载的base64数据
			function() {
				const url = '../../../../_doc/' + new Date() + '.png'; // 生成保存图片的路径，建议使用时间戳命名方式
				console.log('url:', url);
				bitmap.save(
					url, // 保存图片的路径
					{
						overwrite: true // 是否覆盖同名文件
						// quality: 'quality'  // 图片清晰度，此处未启用
					},
					function(i) {
						uni.saveImageToPhotosAlbum({
							filePath: url, // 保存成功后要保存到相册的图片路径
							success: function() {
								uni.showToast({
									title: '保存成功',
									icon: 'none'
								});
								bitmap.clear(); // 清除Bitmap对象占用的内存
							}
						});
					},
					function(e) {
						uni.showToast({
							title: '保存失败',
							icon: 'none'
						});
						console.log(e);
						bitmap.clear();
					}
				);
			},
			function(e) {
				console.log('保存失败', e);
				bitmap.clear();
			}
		);
		return url; // 注意：这里的url变量在异步操作中定义，无法直接返回正确的结果
	},

	dataURItoBlob: function(dataURI) {
		var mimeString = dataURI.split(',')[0].split(':')[1].split(';')[0]; // 获取mime类型
		var byteString = atob(dataURI.split(',')[1]); // 对base64数据进行解码
		var arrayBuffer = new ArrayBuffer(byteString.length); // 创建缓冲数组
		var intArray = new Uint8Array(arrayBuffer); // 创建无符号8位整数视图

		for (var i = 0; i < byteString.length; i++) {
			intArray[i] = byteString.charCodeAt(i); // 将解码后的数据存入视图
		}

		return new Blob([intArray], {
			type: mimeString
		}); // 根据视图创建Blob对象并指定mime类型
	}

}
