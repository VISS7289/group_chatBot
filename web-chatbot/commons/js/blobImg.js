export default{
	//转换时间
	saveBase64Img: function(base64) {
		const bitmap = new plus.nativeObj.Bitmap('test');
		bitmap.loadBase64Data(
			base64,
			function() {
				const url = '../../../../_doc/' + new Date() + '.png'; // url建议用时间戳命名方式
				console.log('url:', url);
				bitmap.save(
					url, {
						overwrite: true // 是否覆盖
						// quality: 'quality'  // 图片清晰度
					},
					i => {
						uni.saveImageToPhotosAlbum({
							filePath: url,
							success: function() {
								uni.showToast({
									title: '保存成功',
									icon: 'none'
								})
								bitmap.clear();
							}
						});
					},
					e => {
						uni.showToast({
							title: '保存失败',
							icon: 'none'
						})
						console.log(e);
						bitmap.clear();
					}
				);
			},
			e => {
				console.log('保存失败', e);
				bitmap.clear();
			}
		);
		return url;
	},
	dataURItoBlob: function(dataURI) {
	        var mimeString = dataURI.split(',')[0].split(':')[1].split(';')[0]; // mime类型
	        var byteString = atob(dataURI.split(',')[1]); //base64 解码
	        var arrayBuffer = new ArrayBuffer(byteString.length); //创建缓冲数组
	        var intArray = new Uint8Array(arrayBuffer); //创建视图
	
	        for (var i = 0; i < byteString.length; i++) {
	            intArray[i] = byteString.charCodeAt(i);
	        }
	        return new Blob([intArray], {type: mimeString});
	}
}