<template>
	<view>
		<view class="submit" @tap="setImage">Test</view>
		<view>
			<canvas style="width: 600px; height: 600px;" canvas-id="cv" type="2d"></canvas>

		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {}
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
			setImage: function() {
				//2. 获取2D上下文
				var ctx = uni.createCanvasContext('cv') 
				//3. 新建一个Image对象
				var img = new Image()
				//4. 设置Image的src
				img.src = '../../static/test/1 (1).jpg'
				//5. 确定Image加载完毕后将Image画到canvas上
				img.onload = () => {
					ctx.drawImage(img, 0, 0, 400, 400)
				}

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
