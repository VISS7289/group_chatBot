<template>
	<view class="content">
		<view class="top-bar">
			<view class="top-bar-left">
				<image src="../../static/general/return.png" class="back-img"></image>
			</view>
			<view class="top-bar-center">
				<view class="title">群聊</view>
			</view>
			<view class="top-bar-right">
				<view class="pice"></view>
				<view class="group-img">
					<image src="../../static/test/1 (17).jpg" class="group-img"></image>
				</view>
			</view>
		</view>
		<scroll-view class="chat" scroll-y="true" scroll-with-animation="true" :scroll-into-view="scrollToView">
			<view class="chat-main" :style="{paddingBottom : h + 'px'}">
				<view class="chat-ls" v-for="(item,index) in msgs" :key="index" :id="'msg'+item.tip">
					<view class="chat-time" v-if="item.time != ''">{{changeTime(item.time)}}</view>
					<view class="msg msg-left" v-if="item.id !='1232'">
						<image :src="item.imgurl" class="user-img"></image>
						<view class="msg-msg" v-if="item.types==0">
							<view class="msg-text">{{item.message}}</view>
						</view>
						<view class="msg-msg" v-if="item.types==1">
							<image :src="item.message" class="msg-img" mode="widthFix" @tap="previewImg(item.message)"></image>
						</view>
					</view>
					<view class="msg msg-right" v-if="item.id =='1232'">
						<image :src="item.imgurl" class="user-img"></image>
						<view class="msg-msg" v-if="item.types==0">
							<view class="msg-text">{{item.message}}</view>
						</view>
						<view class="msg-msg" v-if="item.types==1">
							<image :src="item.message" class="msg-img" mode="widthFix" @tap="previewImg(item.message)"></image>
						</view>
					</view>
				</view>
			</view>
			<view class="padbt"></view>
		</scroll-view>
		<submit @inputChatMsg = "inputChat" @heights="heights"></submit>
	</view>
</template>

<script>
	import datas from '../../commons/js/datas.js';
	import calT from '../../commons/js/calTime.js';
	import submit from '../../components/submit/submit.vue'
	
	export default {
		data() {
			return {
				msgs:[],
				imgMsg:[],
				nowTime: new Date(),
				scrollToView:'',
				h:'120',
			}
		},
		onLoad() {
			this.getMsg();
			this.timer = setInterval(() => {
			    //this.test()
				let len = this.msgs.length;
				let newMsg = this.msgs[len-1];
				if(len == 1){
					this.answerRandomUser(newMsg.message);
				}else{
					let newMsg2 = this.msgs[len-2];
					console.log(newMsg2.message+' '+newMsg.message);
					this.answerRandomUser(newMsg2.message+' '+newMsg.message);
				}
				
				//console.log(this.answerRandomUser(newMsg.message));
			}, 2000+Math.round(4000*Math.random()));
		},
		components: {
			submit,
		},
		methods: {
			//处理时间
			changeTime:function(date){
				return calT.chatTime(date);
			},
			//获取聊天数据
			getMsg: function(){
				let msg = datas.message();
				for(var i=0;i<msg.length;i++){
					msg[i].imgurl='../../static/index/'+msg[i].imgurl;
					//时间间隔
					let t = calT.spaceTime(msg[i].time, this.nowTime);
					if(t) {
						this.nowTime = t;
					}
					msg[i].time = t;
					//信息为图片
					if(msg[i].types==1){
						msg[i].message='../../static/index/'+msg[i].message;
						this.imgMsg.unshift(msg[i].message);
					}
					this.msgs.unshift(msg[i]);
				}
				//滑到当前聊天
				this.$nextTick(function(){
					this.scrollToView='msg'+this.msgs[i-1].tip;
				})
				//console.log(msg);
			},
			//预览图片
			previewImg: function(urlImg) {
				uni.previewImage({
					current: urlImg,
					urls: this.imgMsg,
					longPressActions: {
						itemList: ['发送给朋友', '保存图片', '收藏'],
						success: function(data) {
							console.log('选中了第' + (data.tapIndex + 1) + '个按钮,第' + (data.index + 1) + '张图片');
						},
						fail: function(err) {
							console.log(err.errMsg);
						}
					}
				});
			},
			inputChat: function(inf) {
				let len = this.msgs.length;
				inf=inf.slice(0,inf.length-1)
				let newMsg = {
					id: '1232',
					imgurl: '1 (81).jpg',
					message: inf,
					types: 0,
					time: new Date(),
					tip: len,
				};
				newMsg.imgurl='../../static/index/'+newMsg.imgurl;
				let t = calT.spaceTime(newMsg.time, this.nowTime);
				if(t) {
					this.nowTime = t;
				}
				newMsg.time = t;
				this.msgs.push(newMsg);
				this.goBottom();
				
				this.answerRandomUser(newMsg.message);
				
			},
			answerRandomUser: async function(msg) {
				let retMsg= await this.chat(msg);
				let len = this.msgs.length;
				let theNum = Math.round(145*Math.random()).toString();
				//console.log(retMsg);
				let newMsg = {
					id: '1231',
					imgurl: '1 ('+ theNum +').jpg',
					message: retMsg,
					types: 0,
					time: new Date(),
					tip: len,
				};
				newMsg.imgurl='../../static/index/'+newMsg.imgurl;
				let t = calT.spaceTime(newMsg.time, this.nowTime);
				if(t) {
					this.nowTime = t;
				}
				newMsg.time = t;
				this.msgs.push(newMsg);
				this.goBottom();
			},
			heights: function(inf) {
				this.h = inf;
				this.goBottom();
				console.log(inf);
			},
			//滚动到底部
			goBottom: function() {
				this.scrollToView = '';
				this.$nextTick(function(){
					let len = this.msgs.length - 1;
					this.scrollToView='msg'+this.msgs[len].tip;
				})
			},
			//发送聊天内容
			chat: function(message) {
				return new Promise((resolve, reject) => {
					let retMsg = "1";
					uni.request({
						url: '/chat/chat',
						method: "POST",
						data: JSON.stringify({
							chat_content: message,
						}),
						success: res => {
							retMsg = res.data.Data;
							// console.log(res.data.Data);
							resolve(retMsg)
						},
						fail: () => {
							//reject
						},
						complete: res => {},
					})

				})
				
			},
			test: function() {
				console.log(2);
			}
			
			}
		}

</script>

<style lang="scss">
	@import "../../commons/css/mycss.scss";
	page {
		height: 100%;
	}
	.content{
		height: 100%;
		background: rgba(244,244,244,0.98);
	}
	.top-bar{
		background: rgba(244,244,244,0.98);
		border-bottom: 1px solid $uni-border-color;
		.group-img{
			//float: right;
			bottom: 10rpx;
			width: 68rpx;
			height: 68rpx;
			margin-top: 10rpx;
			margin-right: 10rpx;
			image{
				width: 68rpx;
				height: 68rpx;
				//display: inline-block;
				border-radius: $uni-border-radius-base;
			}
		}
	}
	.chat{
		height: 100%;
		.padbt{
			padding-top: var(--status-bar-height);
		}
		.chat-main{
			padding-left: $uni-spacing-col-base;
			padding-right: $uni-spacing-col-base;
			padding-top: 100rpx;
			//padding-bottom: 120rpx;
			display: flex;
			flex-direction: column;
		}
		.chat-ls{
			.chat-time{
				font-size: $uni-font-size-sm;
				color: rgba(39,40,50,0.3);
				line-height: 34rpx;
				padding: 20rpx 0;
				text-align: center;
			}
			.msg{
				display: flex;
				padding: 20rpx 0;
				.user-img{
					flex: none;
					width: $uni-img-size-sm;
					height: $uni-img-size-sm;
					border-radius: $uni-border-radius-base;
				}
				.msg-msg{
					flex: none;
					max-width: 480rpx;
				}
				.msg-text{
					//width: fit-content;
					font-size: $uni-font-size-lg;
					color: $uni-text-color;
					line-height: 44rpx;
					padding: 18rpx 24rpx;
				}
				.msg-img{
					max-width: 464rpx;
					border-radius: 20rpx;
				}
			}
			.msg-left{
				flex-direction: row;
				.msg-text{
					margin-left: 16rpx;
					background-color: #fff;
					border-radius: 0rpx 20rpx 20rpx 20rpx;
				}
				.msg-img{
					margin-left: 16rpx;
				}
			}
			.msg-right{
				flex-direction: row-reverse;
				.msg-text{
					margin-right: 16rpx;
					background-color: rgba(255,228,49,0.8);
					border-radius: 20rpx 0rpx 20rpx 20rpx;
				}
				.msg-img{
					margin-right: 16rpx;
				}
			}
		}
	}
	
</style>
