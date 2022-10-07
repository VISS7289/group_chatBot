<template>
	<view class="content">
		<view class="top-bar">
			<view class="top-bar-left" @tap="backOne">
				<image src="../../static/general/return.png" class="back-img"></image>
			</view>
			<view class="top-bar-center">
				<view class="title">{{friend.name}}</view>
			</view>
			<view class="top-bar-right">
				<view class="group-img">
					<image :src="friend.img" mode="aspectFill"></image>
				</view>
			</view>
		</view>
		<scroll-view class="chat" scroll-y="true" scroll-with-animation="true" :scroll-into-view="scrollToView">
			<view class="chat-main" :style="{paddingBottom : h + 'px'}">
				<view class="chat-ls" v-for="(item,index) in msgs" :key="index" :id="'msg'+item.tip">
					<view class="chat-time" v-if="changeTime3(index)!=''">{{changeTime3(index)}}</view>
					<view class="msg msg-left" v-if="item.id != user.id">
						<image :src="item.img" class="user-img" mode="aspectFill"></image>
						<view class="msg-msg" v-if="item.type==0">
							<view class="msg-text">{{item.msg}}</view>
						</view>
						<view class="msg-msg" v-if="item.type==1">
							<image :src="item.msg" class="msg-img" mode="widthFix" @tap="previewImg(item.msg)">
							</image>
						</view>
					</view>
					<view class="msg msg-right" v-if="item.id == user.id">
						<image :src="user.img" class="user-img" mode="aspectFill"></image>
						<view class="msg-msg" v-if="item.type==0">
							<view class="msg-text">{{item.msg}}</view>
						</view>
						<view class="msg-msg" v-if="item.type==1">
							<image :src="item.msg" class="msg-img" mode="widthFix" @tap="previewImg(item.msg)">
							</image>
						</view>
					</view>
				</view>
			</view>
			<view class="padbt"></view>
		</scroll-view>
		<submit @inputChatMsg="inputChat" @heights="heights"></submit>
	</view>
</template>

<script>
	import datas from '../../commons/js/datas.js'
	import calT from '../../commons/js/calTime.js'
	import submit from '../../components/submit/submit.vue'
	import config from '../../commons/js/config.js'
	import refersh from '../../commons/js/refershToken.js'
	export default {
		data() {
			return {
				msgs: [],
				imgMsg: [],
				nowTime: calT.getNewTime(),
				scrollToView: '',
				h: '120',
				friend: {},
				atoken: '',
				rtoken: '',
				user: {},
				nowpage: 0,
				maxpage: 10,
			}
		},
		onLoad(option) {
			this.userInit()
			this.friend = JSON.parse(decodeURIComponent(option.friendInfo))
			console.log(this.friend)
			console.log(option.type)
			this.getMsg2()
			// this.timer = setInterval(() => {
			//     //this.test()
			// 	let len = this.msgs.length;
			// 	let newMsg = this.msgs[len-1];
			// 	if(len == 1){
			// 		this.answerRandomUser(newMsg.message);
			// 	}else{
			// 		let newMsg2 = this.msgs[len-2];
			// 		//console.log(newMsg2.message+'.'+newMsg.message);
			// 		console.log(3);
			// 		this.answerRandomUser(newMsg2.message+' '+newMsg.message);
			// 	}

			// 	//console.log(this.answerRandomUser(newMsg.message));
			// }, 2000+Math.round(4000*Math.random()));
		},
		components: { submit, },
		methods: {
			userInit: function() {
				try {
					const value = uni.getStorageSync('user')
					if (value) {
						this.user.img = 'data:image/png;base64,' + value.img
						this.user.id = value.id
						this.user.email = value.email
						this.user.name = value.name
						this.user.nick = value.name
						this.atoken = uni.getStorageSync('atoken')
						this.rtoken = uni.getStorageSync('rtoken')
					} else {
						uni.navigateTo({ url: '../signin/signin', })
					}
					console.log(value)
				} catch (e) {
					// error
				}
			},
			backOne: function() {
				uni.navigateBack({ delta: 1 })
			},
			//处理时间
			changeTime: function(date) {
				return calT.chatTime(date)
			},
			changeTime2: function(date) {
				return calT.dateTime2(date)
			},
			changeTime3: function(index) {
				if(index==0){
					return calT.dateTime2(this.msgs[index].time)
				}
				let delta=calT.spaceTime2(this.msgs[index-1].time,this.msgs[index].time)
				let equal=(calT.dateTime2(this.msgs[index-1].time)==calT.dateTime2(this.msgs[index].time))
				console.log(delta)
				if(delta>5 && equal == false){
					return calT.dateTime2(this.msgs[index].time)
				}else{
					return ''
				}
			},
			//获取聊天数据
			getMsg: function() {
				let msg = datas.message()
				for (var i = 0; i < msg.length; i++) {
					msg[i].imgurl = '../../static/index/' + msg[i].imgurl
					//时间间隔
					let t = calT.spaceTime(msg[i].time, this.nowTime)
					if (t) {
						this.nowTime = t
					}
					msg[i].time = t
					//信息为图片
					if (msg[i].types == 1) {
						msg[i].message = '../../static/index/' + msg[i].message
						this.imgMsg.unshift(msg[i].message)
					}
					this.msgs.unshift(msg[i])
				}
				//滑到当前聊天
				this.$nextTick(function() {
					this.scrollToView = 'msg' + this.msgs[i - 1].tip
				})
				//console.log(msg);
			},
			getMsg2: function() {
				if (this.nowpage == -1) {
					return
				}
				uni.request({
					url: config.myurl + '/msg/getOldChatF',
					method: 'POST',
					header: { 'Authorization': 'Bearer ' + this.atoken },
					data: {
						'user_id': this.user.id,
						'friend_id': this.friend.id,
						'now_page': this.nowpage,
						'max_page': this.maxpage,
					},
					success: async data => {
						console.log(data.data)
						if (data.data.Code == 1009) {
							let newCode = await refersh.refersh(config.myurl, this.atoken, this
								.rtoken)
							if (newCode == 1000) {
								this.atoken = uni.getStorageSync('atoken')
								this.rtoken = uni.getStorageSync('rtoken')
								this.getMsg2()
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
							console.log(data.data.Data)
							let l = this.msgs.length
							for (var i = 0; i < data.data.Data.length; i++) {
								console.log(data.data.Data[i].Time)
								console.log(this.nowTime)
								// let t = calT.spaceTime(data.data.Data[i].Time, this.nowTime)
								// if (t) {
								// 	this.nowTime = t
								// }
								
								this.msgs.push({
									id: data.data.Data[i].SendId,
									img: 'data:image/png;base64,' + data.data.Data[i].Img,
									msg: data.data.Data[i].Msg,
									time: data.data.Data[i].Time,
									type: data.data.Data[i].Type,
									tip: l + i
								})
							}
							//滑到当前聊天
							this.$nextTick(function() {
								this.scrollToView = 'msg' + this.msgs[data.data.Data.length - 1]
									.tip
							})
							if (data.data.Data.length == 10) {
								this.nowpage++
							} else {
								//数据提取完毕
								this.nowpage = -1
							}
						} else {
							console.log('err')
						}
					}
				})
			},
			//预览图片
			previewImg: function(urlImg) {
				uni.previewImage({
					current: urlImg,
					urls: this.imgMsg,
					longPressActions: {
						itemList: ['发送给朋友', '保存图片', '收藏'],
						success: function(data) {
							console.log('选中了第' + (data.tapIndex + 1) + '个按钮,第' + (data.index + 1) + '张图片')
						},
						fail: function(err) {
							console.log(err.errMsg)
						}
					}
				})
			},
			inputChat: function(inf) {
				console.log(calT.getNewTime())
				this.reciveMsg(this.user.id,this.user.img,inf,0)
			},
			reciveMsg: function(id,img,msg,type) {
				let len = this.msgs.length
				let newMsg = {
					id: id,
					img: img,
					msg: msg,
					type: type,
					time: calT.getNewTime(),
					tip: len,
				}
				// let t = calT.spaceTime2(this.nowTime, newMsg.time)
				// if (t) {
				// 	this.nowTime = t
				// }
				// newMsg.time = t
				console.log(this.changeTime2(newMsg.time))
				this.msgs.push(newMsg)
				this.goBottom()
			
			},
			answerRandomUser: async function(msg) {

				let retMsg = await this.chat(msg)
				let len = this.msgs.length
				let theNum = Math.round(145 * Math.random()).toString()

				console.log(retMsg)
				let newMsg = {
					id: '1231',
					imgurl: '1 (' + theNum + ').jpg',
					message: retMsg,
					types: 0,
					time: new Date(),
					tip: len,
				}
				newMsg.imgurl = '../../static/index/' + newMsg.imgurl
				let t = calT.spaceTime(newMsg.time, this.nowTime)
				if (t) {
					this.nowTime = t
				}
				newMsg.time = t
				this.msgs.push(newMsg)
				this.goBottom()
			},
			heights: function(inf) {
				this.h = inf
				this.goBottom()
				console.log(inf)
			},
			//滚动到底部
			goBottom: function() {
				this.scrollToView = ''
				this.$nextTick(function() {
					let len = this.msgs.length - 1
					this.scrollToView = 'msg' + this.msgs[len].tip
				})
			},
			//发送聊天内容
			chat: function(message) {
				return new Promise((resolve, reject) => {
					let retMsg = '1'

					const xhr = new XMLHttpRequest()
					const url = 'http://localhost:8086/ping'

					xhr.open('GET', url)
					xhr.onreadystatechange = function() {
						console.log(xhr.responseText)
					}
					xhr.send()


					uni.request({
						url: 'http://localhost:8086/chat', //仅为示例，并非真实接口地址。
						method: 'POST',
						data: JSON.stringify({ chat_content: message, }),
						header: { 'custom-header': 'hello' //自定义请求头信息
						},
						success: res => {
							retMsg = res.data.Data
							resolve(retMsg)
						},
					})
					// uni.request({
					// 	//url: '/chat/chat',
					// 	url: '/chat/chat',
					// 	method: "POST",
					// 	data: JSON.stringify({
					// 		chat_content: message,
					// 	}),
					// 	sslVerify:  false,
					// 	success: res => {
					// 		retMsg = res.data.Data;
					// 		resolve(retMsg)
					// 	},
					// 	fail: res => {
					// 		console.log(0);
					// 		console.log(res);
					// 		//reject
					// 	},
					// 	complete: res => {
					// 		console.log(2);
					// 	},
					// })

				})

			},
			test: function() {
				//console.log(2);
			}

		}
	}
</script>

<style lang="scss">
	@import "../../commons/css/mycss.scss";

	page {
		height: 100%;
	}

	.content {
		height: 100%;
		background: rgba(244, 244, 244, 0.98);
	}

	.top-bar {
		background: rgba(244, 244, 244, 0.98);
		border-bottom: 1px solid $uni-border-color;

		.top-bar-right {
			margin-right: 9rpx;

			.group-img {
				//float: right;
				// bottom: 10rpx;
				width: 68rpx;
				height: 68rpx;
				// margin-top: 10rpx;
				// margin-right: 10rpx;
				margin: 0;
				padding: 0rpx;

				image {
					margin: 9rpx 0;
					width: 68rpx;
					height: 68rpx;
					//display: inline-block;
					border-radius: $uni-border-radius-base;
				}
			}
		}

	}

	.chat {
		height: 100%;

		.padbt {
			padding-top: var(--status-bar-height);
		}

		.chat-main {
			padding-left: $uni-spacing-col-base;
			padding-right: $uni-spacing-col-base;
			padding-top: 100rpx;
			//padding-bottom: 120rpx;
			display: flex;
			flex-direction: column;
		}

		.chat-ls {
			.chat-time {
				font-size: $uni-font-size-sm;
				color: rgba(39, 40, 50, 0.3);
				line-height: 34rpx;
				padding: 20rpx 0;
				text-align: center;
			}

			.msg {
				display: flex;
				padding: 20rpx 0;

				.user-img {
					flex: none;
					width: $uni-img-size-sm;
					height: $uni-img-size-sm;
					border-radius: $uni-border-radius-base;
				}

				.msg-msg {
					flex: none;
					max-width: 480rpx;
				}

				.msg-text {
					//width: fit-content;
					font-size: $uni-font-size-lg;
					color: $uni-text-color;
					line-height: 44rpx;
					padding: 18rpx 24rpx;
				}

				.msg-img {
					max-width: 464rpx;
					border-radius: 20rpx;
				}
			}

			.msg-left {
				flex-direction: row;

				.msg-text {
					margin-left: 16rpx;
					background-color: #fff;
					border-radius: 0rpx 20rpx 20rpx 20rpx;
				}

				.msg-img {
					margin-left: 16rpx;
				}
			}

			.msg-right {
				flex-direction: row-reverse;

				.msg-text {
					margin-right: 16rpx;
					background-color: rgba(255, 228, 49, 0.8);
					border-radius: 20rpx 0rpx 20rpx 20rpx;
				}

				.msg-img {
					margin-right: 16rpx;
				}
			}
		}
	}
</style>
