<template>
	<view class="content">
		<view class="top-bar">
			<view class="top-bar-right" @tap="toSignUp">
				<view class="text">注册</view>
			</view>
		</view>
		<view class="logo">
			<image src='../../static/general/logo.png'></image>
		</view>
		<view class="main">
			<view class="title">登录</view>
			<view class="slogan">您好，欢迎来到AI聊天室！</view>
			<view class="inputs">
				<input type="text" placeholder="用户名/邮箱" class="user" placeholder-style="color:#bbb;font-weight:400;" @blur="subUsername"/>
				<input type="password" placeholder="密码" class="paw" placeholder-style="color:#bbb;font-weight:400;" @blur="subPassword"/>
			</view>
			<view class="tips" v-if="wrong">输入用户名或密码错误！</view>
		</view>
		<view class="submit" @tap="subInfo">登录</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				wrong: false,
				username: '',
				password: '',
			}
		},
		methods: {
			toSignUp: function(){
				uni.navigateTo({
					url: '../signup/signup',
				});
			},
			subUsername: function(e){
				this.username=e.detail.value;
			},
			subPassword: function(e){
				this.password=e.detail.value;
			},
			subInfo: function(){
				if(this.username && this.password){
					this.wrong=false;
					console.log(this.username);
					console.log(this.password);
				}else{
					this.wrong=true;
				}
			}
		}
	}
</script>

<style lang="scss">
	@import "../../commons/css/mycss.scss";
	.top-bar{
		border-bottom: 1px solid $uni-border-color;
		.top-bar-right{
			image{
				padding: 18rpx 0 0 18rpx;
				width: 52rpx;
				height: 52rpx;
			}
			.text{
				font-size: 40rpx;
				font-weight: 500;
				color: $uni-text-color;
				line-height: 80rpx;
			}
		}
	}
	.logo{
		text-align: center;
		image{
		padding-top: 192rpx;
		width: 194rpx;
		height: 100rpx;
		}
	}
	.main{
		padding:54rpx $uni-spacing-row-lg 120rpx;
		height: 458rpx;
		.title{
			font-size: 56rpx;
			font-weight: 500;
			color: $uni-text-color;
			line-height: 80rpx;
		}
		.slogan{
			font-size: 40rpx;
			color:$uni-text-color-grey;
			line-height: 56rpx;
		}
		.inputs{
			padding-top: 8rpx;
			input{
				padding-top: 40rpx;
				height: 88rpx;
				font-size: $uni-font-size-lg;
				font-weight: 500;
				color: $uni-text-color;
				line-height: 88rpx;
				border-bottom: 1px solid $uni-border-color;
			}
		}
		.tips{
			float: left;
			font-size: $uni-font-size-lg;
			color: $uni-color-warning;
			line-height: 56rpx;
		}
	}
	.submit{
		margin: 0 auto;
		width: 520rpx;
		height: 96rpx;
		background: $uni-color-primary;
		box-shadow: 0rpx 50rpx 32rpx -36rpx rgba(255,228,49,0.4);
		border-radius: 48rpx;
		font-size: $uni-font-size-lg;
		font-weight: 500;
		color: $uni-text-color;
		line-height: 96rpx;
		text-align: center;
	}
</style>
