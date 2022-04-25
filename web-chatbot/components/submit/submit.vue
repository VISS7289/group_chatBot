<template>
	<view>
		<view class="submit">
			<view class="submit-chat">
				<view class="bt-img" @tap="records">
					<image :src="toc"></image>
				</view>
				<textarea auto-height="true" class="chat-send btn" v-show="!isrecord" @input="inputs" v-model="msg"></textarea>
				<view class="record btn" v-show="isrecord">按住说话</view>
				<view class="bt-img" @tap="emoji">
					<image src="../../static/general/smile.png"></image>
				</view>
				<view class="bt-img">
					<image src= "../../static/general/add.png"></image>
				</view>
			</view>
			<view class="emoji" v-show="isemoji">
				表情
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				isrecord: false,
				isemoji: false,
				msg: '',
				toc:"../../static/general/record.png"
			}
		},
		methods: {
			//获取高度数据
			getHeight: function() {
				const query = uni.createSelectorQuery().in(this);
				query.select('.submit').boundingClientRect(data=>{
					this.$emit('heights',data.height);
				}).exec();
			},
			//切换音频输入
			records: function(){
				if(this.isrecord==true){
					this.isrecord=false;
					this.toc="../../static/general/record.png";
				}else{
					this.isrecord=true;
					this.toc="../../static/general/keybord.png";
				}
				
			},
			//表情
			emoji: function(){
				this.isemoji=!this.isemoji;
				setTimeout(()=>{
					this.getHeight();
				},0);
			},
			//文字发送
			inputs: function(inf) {
				var chat = inf.detail.value;
				var pos = chat.indexOf('\n');
				if(pos != -1 && chat.length > 1) {
					this.$emit('inputChatMsg',this.msg);
					setTimeout(()=>{
						this.msg = '';
					},0);
				}
			}
		}
	}
</script>

<style lang="scss">
	.submit{
		background: rgba(244,244,244,0.96);
		border-top: 1px solid $uni-border-color;
		width: 100%;
		position: fixed;
		bottom: 0;
		left: 0;
		z-index: 100;
		padding-bottom: var(--status-bar-height);
	}
	.submit-chat{
		width: 100%;
		//height: 520rpx;
		display: flex;
		align-items: flex-end;
		box-sizing: border-box;
		padding: 14rpx 14rpx;
		.bt-img{
			
		}
		.btn{
			flex: auto;
			background-color: #fff;
			border-radius: 10rpx;
			padding: 20rpx;
			margin: 0 10rpx;
			max-height: 160rpx;
			font-size: $uni-font-size-lg;
		}
		.record{
			text-align: center;
			
			color: $uni-text-color-grey;
		}
		image{
			width: 56rpx;
			height: 56rpx;
			margin: 0 10rpx;
			flex: auto;
		}
	}
	.emoji{
		height: 460rpx;
		width: 100%;
		background: rgba(236,237,238,1);
		box-shadow: 0rpx -1rpx 0rpx 0rpx rgba(0,0,0,0.1);
	}
</style>
