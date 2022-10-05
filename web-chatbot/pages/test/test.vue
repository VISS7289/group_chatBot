<template>
	<view>
		<view class="submit" @tap="setImage2">Test</view>
		<view>
			<canvas style="width: 100px; height: 100px; position:fixed;left:100%;" canvas-id="cv" type="2d"></canvas>
		</view>
		<image :src="myimg"></image>
	</view>
</template>

<script>
	export default {
		data() {
			return { myimg: '', }
		},

		methods: {
			Test: function() {
				var img = document.createElement('img')
				var tempP = ''
				img.src = '../../static/test/1 (1).jpg'
				uni.compressImage({
					src: img.src,
					quality: 80,
					success: res => {
						tempP = res.tempFilePath
						console.log(res.tempFilePath)
					}
				})
				uni.request({
					url: tempP,
					method: 'GET',
					responseType: 'arraybuffer',
					success: ress => {
						let base64 = wx.arrayBufferToBase64(ress.data) //把arraybuffer转成base64 
						base64 = 'data:image/jpeg;base64,' + base64 //不加上这串字符，在页面无法显示的哦
						console.log(base64)
					}
				})
			},
			getImgInfo(url) {
				return new Promise((resolve, reject) => {
					let image = new Image()
					image.src = url
					let timer = setInterval(() => {
						if (image.width > 0 || image.height > 0) {
							resolve([image.width, image.height]) // 图片宽*高
							clearInterval(timer)
						}
					}, 50)
				})
			},
			setImage: async function() {
				//2. 获取2D上下文
				var ctx = uni.createCanvasContext('cv',this)
				//4. 设置Image的src
				let src = '../../static/test/1 (20).jpg'
				let w, h
				//5. 确定Image加载完毕后将Image画到canvas上
				await this.getImgInfo(src).then(res => {
					w = res[0]
					h = res[1]
				})
				let d = this.Afill(w, h, 100, 100)
				console.log()
				ctx.drawImage(src, d[0], d[1], d[2], d[3], 0, 0, 100, 100)
				ctx.draw(false,()=>{
					uni.canvasToTempFilePath({
					  canvasId: 'cv',
					  fileType: 'jpg',
					  quality: 0.8,
					  success: function(res) {
					    // 在H5平台下，tempFilePath 为 base64
					    console.log(res.tempFilePath)
						this.myimg=res.tempFilePath
						console.log(this.myimg)
					  } 
					},this)
				}) //绘制
				
			},
			setImage2: async function() {
				let src = '../../static/test/1 (20).jpg'
				let res= await this.compressImg(src)
				console.log(res)
				
			},
			compressImg: function(src){
				return new Promise(async (resolve, reject)=> {
					var ctx = uni.createCanvasContext('cv',this)
					let w, h
					await this.getImgInfo(src).then(res => {
						w = res[0]
						h = res[1]
					})
					let d = this.Afill(w, h, 100, 100)
					console.log()
					ctx.drawImage(src, d[0], d[1], d[2], d[3], 0, 0, 100, 100)
					ctx.draw(false,()=>{
						uni.canvasToTempFilePath({
						  canvasId: 'cv',
						  fileType: 'jpg',
						  quality: 0.8,
						  success: function(res) {
						    // 在H5平台下，tempFilePath 为 base64
							console.log(res.tempFilePath)
							resolve(res.tempFilePath)
						  } ,
						  fail: function() {
							  reject('error')
						  }
						},this)
					}) //绘制
				}) 
			},
			Afill: function(imageWidth, imageHeight, canvasWidth, canvasHeight) {
				const imageRate = imageWidth / imageHeight
				const canvasRate = canvasWidth / canvasHeight
				let [sx, sy, sw, sh] = []
				if (imageRate >= canvasRate) {
					sw = imageHeight * canvasRate
					sh = imageHeight
					sx = (imageWidth - sw) / 2
					sy = 0
				} else {
					sh = imageWidth / canvasRate
					sw = imageWidth
					sx = 0
					sy = (imageHeight - sh) / 2
				}
				return [sx, sy, sw, sh]
			},


		}
	}
</script>

<style>
	@import "../../commons/css/mycss.scss";

	.submit {
		margin: 0 auto;
		width: 520rpx;
		height: 96rpx;
		background: $uni-color-primary;
		box-shadow: 0rpx 50rpx 32rpx -36rpx rgba(255, 228, 49, 0.4);
		border-radius: 48rpx;
		font-size: $uni-font-size-lg;
		font-weight: 500;
		color: $uni-text-color;
		line-height: 96rpx;
		text-align: center;
	}
</style>
