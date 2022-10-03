<template>
	<view class="content">
		<view class="top-bar">
			<view class="serch-div">
				<input type="serch" placeholder="搜索用户/群" class="serch" placeholder-style="color:#bbb;font-weight:400;"
					@input="searchKey" />
				<image src="../../static/index/search.png" class="serch-img"></image>
			</view>
			<view class="top-bar-right" @tap="toIndex">
				<view class="text">取消</view>
			</view>
		</view>
		<view class="main">
			<view class="serchUser">
				<view class="titleUser" v-if="userArr.length>0">用户</view>
				<view class="listUser" v-for="(item,index) in userArr" :key="index">
					<navigator :url="'../userhome/userhome?user='+encodeURIComponent(JSON.stringify(item))"
						hover-class="none">
						<image :src="item.img" mode="aspectFill"></image>
					</navigator>
					<view class="nameUser" v-html="item.username"></view>
					<view class="rightBt info" v-if="item.state==0">发消息</view>
					<view class="rightBt friend" v-if="item.state==1" @tap="goModify(6,item,index)">申请中</view>
					<view class="rightBt friend" v-if="item.state==2" @tap="goModify(6,item,index)">加好友</view>
				</view>
			</view>
			<!-- 			<view class="serchUser">
				<view class="titleUser">群组</view>
				<view class="listUser">
					<image src="../../static/index/1 (1).jpg"></image>
					<view class="nameUser">实力至上</view>
					<view class="rightBt info">发消息</view>
					<view class="rightBt friend" v-if="dis">加好友</view>
				</view>
				<view class="listUser">
					<image src="../../static/index/1 (81).jpg"></image>
					<view class="nameUser">大力出奇迹</view>
					<view class="rightBt info">发消息</view>
					<view class="rightBt friend" v-if="dis">加好友</view>
				</view>
			</view> -->
		</view>
	</view>
</template>

<script>
	import datas from '../../commons/js/datas.js'
	import Debounce from '../../commons/js/delay.js'
	import config from '../../commons/js/config.js'
	import refersh from '../../commons/js/refershToken.js'
	export default {
		data() {
			return {
				dis: false,
				userArr: [],
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
					type: 6, //infoType: ['昵称','备注','简介','邮箱','电话','加好友'],
					userRequest: '请求加为好友'
				},
				atoken: '',
				rtoken: '',
				queryIndex: 0,
			}
		},
		onLoad: function(option) {
			this.user = JSON.parse(decodeURIComponent(option.user))
			try {
				this.atoken = uni.getStorageSync('atoken')
				this.rtoken = uni.getStorageSync('rtoken')
			} catch (e) {
				// error
			}
		},
		methods: {
			searchKey: Debounce.Debounce(function(e) {
				this.userArr = []
				let serchVal = e.detail.value
				if (serchVal.length > 0) {
					this.userDel(serchVal)
				}
			}, 500),
			AddSuccess:function(){
				this.userArr[this.queryIndex].state=1
			},
			userDel: function(e) {
				console.log(e)
				uni.request({
					url: config.myurl + '/serch/username',
					method: 'POST',
					header: { 'Authorization': 'Bearer ' + this.atoken },
					data: {
						'username': e,
						'user_id': this.user.id
					},
					success: data => {
						console.log(data.data.Data)
						if (data.data.Code == 1009) {
							console.log(this.atoken)
							let newCode = refersh.refersh(config.myurl, this.atoken, this.rtoken)
							console.log(this.atoken)
							if (newCode == 1000) {
								this.userDel()
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
							//success
							this.userArr = []
							let myArr = data.data.Data
							console.log(data.data.Data)
							let light = eval('/' + e + '/g')
							console.log(data.data)
							for (let i = 0; i < myArr.length; i++) {
								if (myArr[i].Username.search(e) != -1) {
									//0好友 1申请中 2非好友 3自己
									if (myArr[i].UserId == this.user.id) {
										continue
									}
									let item = {
										'img': 'data:image/png;base64,' + myArr[i].Img,
										'name': myArr[i].Username,
										'username': myArr[i].Username.replace(light,
											'<span style=\'color:#4AAAFF;\'>' + e +
											'</span>'),
										'state': myArr[i].State,
										'id': myArr[i].UserId
									}
									this.userArr.push(item)
								}
							}
						} else {
							// err
						}
					}
				})

			},
			toIndex: function() {
				uni.navigateTo({ url: '../index/index', })
			},
			goModify: function(type, item, index) {
				this.queryIndex = index
				this.modifyInfo.type = type
				this.modifyInfo.userid = this.user.id
				this.modifyInfo.friendid = item.id
				this.modifyInfo.userRequest = this.user.name + '请求加为好友~'
				uni.navigateTo({ url: '../modify/modify?modifyInfo=' + encodeURIComponent(JSON.stringify(this.modifyInfo)) })
			}
		}
	}
</script>

<style lang="scss">
	@import "../../commons/css/mycss.scss";

	.top-bar {
		background-color: rgba(255, 255, 255, 0.98);
		border-bottom: 1px solid $uni-border-color;

		.serch-div {
			position: absolute;
			top: 0;
			left: 0;
			width: 100%;
			z-index: -1;
			padding: 14rpx;
			margin-top: var(--status-bar-height);

			.serch {
				width: 540rpx;
				height: 60rpx;
				background-color: $uni-bg-color-grey;
				border-radius: 10rpx;
				padding: 0rpx 50rpx 0rpx 18rpx;
			}

			.serch-img {
				position: absolute;
				right: 165rpx;
				top: 25rpx;
				width: 40rpx;
				height: 40rpx;
			}
		}

		.top-bar-right {
			padding-right: 0rpx;

			.text {
				font-size: 40rpx;
			}
		}
	}

	.main {
		width: 100%;
		padding-top: 88rpx;
		padding-left: 32rpx;

		.serchUser {
			width: 100%;
			padding-top: $uni-spacing-col-base;

			.titleUser {
				font-size: 44rpx;
				font-weight: 600;
				color: $uni-text-color;
				line-height: 60rpx;
			}

			.listUser {
				display: flex;
				flex-direction: row;
				align-items: center;
				//justify-content: center;
				width: 100%;
				margin-top: 40rpx;

				image {
					width: 80rpx;
					height: 80rpx;
					float: left;
					border-radius: $uni-border-radius-base;
				}

				.nameUser {
					width: 500rpx;
					padding-left: 32rpx;
					font-size: 36rpx;
					font-weight: 400rpx;
					line-height: 50rpx;
					color: $uni-text-color;
				}

				.rightBt {
					width: 120rpx;
					height: 60rpx;
					border-radius: $uni-border-radius-base;

					font-size: $uni-font-size-sm;
					line-height: 60rpx;
					text-align: center;
				}

				.info {
					background-color: $uni-color-primary;
					color: $uni-text-color;
				}

				.friend {
					background-color: $uni-bg-color-grey;
					color: $uni-color-success;
				}
			}
		}
	}
</style>
