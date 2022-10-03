<template>
	<view class="content">
		<view class="top-bar">
			<view class="top-bar-left" @tap="backOne">
				<image src="../../static/general/return.png"></image>
			</view>
			<view class="top-bar-center">
				<view class="title">详情</view>
			</view>
		</view>
		<view class="main">
			<view class="line">
				<view class="lineHead">头像</view>
				<view class="lineInfo" @click="select">
					<image :src="user.imgurl" mode="aspectFill" class="lineImg"></image>
				</view>
				<ksp-cropper mode="free" :width="200" :height="140" :maxWidth="1024" :maxHeight="1024" :url="url"
					@cancel="oncancel" @ok="onok"></ksp-cropper>
				<view class="lineBtn">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line">
				<view class="lineHead">签名</view>
				<view class="linetext">{{user.intr}}</view>
				<view class="lineBtn">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line">
				<view class="lineHead">注册</view>
				<view class="linetext">2020-10-17</view>
			</view>
			<view class="line">
				<view class="lineHead">昵称</view>
				<view class="linetext">{{user.name}}</view>
				<view class="lineBtn" @tap="goModify(0)">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line">
				<view class="lineHead">备注</view>
				<view class="linetext">{{user.nick}}</view>
				<view class="lineBtn" @tap="goModify(5)">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line">
				<view class="lineHead">性别</view>
				<view class="linetext">
					<view class="uni-list-cell-db">
						<picker @change="bindPickerChange" :value="user.sex" :range="array">
							<view class="uni-input">{{array[user.sex]}}</view>
						</picker>
					</view>
				</view>
				<view class="lineBtn">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line">
				<view class="lineHead">生日</view>
				<view class="linetext">
					<view class="uni-list-cell-db">
						<picker mode="date" :value="date" :start="startDate" :end="endDate" @change="bindDateChange">
							<view class="uni-input">{{date}}</view>
						</picker>
					</view>
				</view>
				<view class="lineBtn">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line">
				<view class="lineHead">电话</view>
				<view class="linetext">{{user.phone}}</view>
				<view class="lineBtn" @tap="goModify(4)">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line">
				<view class="lineHead">邮箱</view>
				<view class="linetext">{{user.email}}</view>
				<view class="lineBtn" @tap="goModify(3)">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line">
				<view class="lineHead">密码</view>
				<view class="linetext">*********</view>
				<view class="lineBtn">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="btn">
				退出登录
			</view>
		</view>
	</view>
</template>

<script>
	import ksp from '../../commons/js/ksp.js'
	//import {saveBase64Img,dataURItoBlob} from '../../commons/js/blobImg.js';
	export default {
		data() {
			return {
				user: {
					id: 4,
					imgurl: '../../static/index/1 (4).jpg',
					tip: 3,
					name: '文嘉',
					time: new Date(),
					info: '明日复明日，明日何其多！日日待明日，万事成蹉跎。世人皆被明日累，明日无穷老将至。晨昏滚滚水东流。今古悠悠日西坠。百年明日能几何？请君听我《明日歌》。',
					sex: 0,
					nick: '嘉嘉嘉',
					intr: '明日复明日，明日何其多！日日待明日，万事成蹉跎。世人皆被明日累，明日无穷老将至。晨昏滚滚水东流。今古悠悠日西坠。百年明日能几何？请君听我《明日歌》。',
					email: '123456@12.com',
					phone: '12345678',
				},
				modifyInfo: {
					type: 2, //infoType: ['昵称','备注','简介','邮箱','电话'],
					userinfo: {}
				},
				array: ['男', '女', '未知'],
				date: this.getDate({ format: true }),
				url: '',
			}
		},
		onLoad: function(option) {
			this.user = JSON.parse(decodeURIComponent(option.user))
		},
		computed: {
			startDate() {
				return this.getDate('start')
			},
			endDate() {
				return this.getDate('end')
			}
		},
		methods: {
			backOne: function() {
				uni.navigateBack({ delta: 1 })
			},
			bindPickerChange: function(e) {
				this.user.sex = e.detail.value
			},
			bindDateChange: function(e) {
				this.date = e.detail.value
			},
			getDate(type) {
				const date = new Date()
				let year = date.getFullYear()
				let month = date.getMonth() + 1
				let day = date.getDate()

				if (type === 'start') {
					year = year - 60
				} else if (type === 'end') {
					year = year + 2
				}
				month = month > 9 ? month : '0' + month
				day = day > 9 ? day : '0' + day
				return `${year}-${month}-${day}`
			},
			select() {
				uni.chooseImage({
					count: 1,
					success: rst => {
						// 设置url的值，显示控件
						this.url = rst.tempFilePaths[0]
					}
				})
			},
			onok(ev) {
				this.url = ''
				this.user.imgurl = ev.path
				//console.log(this.user.imgurl);
			},
			oncancel() {
				// url设置为空，隐藏控件
				this.url = ''
			},
			goModify: function(type) {
				this.modifyInfo.type = type
				this.modifyInfo.userinfo = this.user
				uni.navigateTo({ url: '../modify/modify?modifyInfo=' + encodeURIComponent(JSON.stringify(this.modifyInfo)) })
			}
		}
	}
</script>

<style lang="scss">
	@import "../../commons/css/mycss.scss";

	.top-bar {
		background: rgba(255, 255, 255, 0.96);
		border-bottom: 1px solid $uni-border-color;
		margin-bottom: 30rpx;
	}

	.main {
		display: flex;
		flex-direction: column;
		align-items: center;
		margin-top: 88rpx;
		padding-top: var(--status-bar-height);
		width: 100%;

		.line {
			margin-top: 12rpx;
			width: 100%;
			height: 112rpx;
			display: flex;
			flex-direction: cow;
			align-items: center;

			.lineHead {
				flex: none;
				padding-left: $uni-spacing-col-base;
				font-size: $uni-font-size-lg;
				color: $uni-text-color;
				line-height: 112rpx;
			}

			.linetext {
				flex: auto;
				font-size: $uni-font-size-lg;
				color: $uni-text-color-grey;
				line-height: 112rpx;
				overflow: hidden;
				text-overflow: ellipsis;
				white-space: nowrap;
				padding: 0rpx $uni-spacing-col-base;
			}

			.lineInfo {
				flex: auto;
				display: flex;
				flex-direction: cow;
				align-items: center;
				padding: 0rpx $uni-spacing-col-base;

				.lineImg {
					width: $uni-img-size-lg;
					height: $uni-img-size-lg;
					border-radius: $uni-border-radius-base;
				}
			}

			.lineBtn {
				flex: none;
				padding-right: $uni-spacing-col-base;

				image {
					width: 50rpx;
					height: 50rpx;
				}
			}
		}

		.btn {
			margin-top: 148rpx;
			font-size: $uni-font-size-lg;
			color: $uni-color-warning;
		}
	}
</style>
