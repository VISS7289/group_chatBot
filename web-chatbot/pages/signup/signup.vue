<template>
	<view class="content">
		<view class="top-bar">
			<view class="top-bar-left" @tap="toSignIn">
				<image src="../../static/general/return.png"></image>
			</view>
		</view>
		<view class="logo">
			<image src='../../static/general/logo.png'></image>
		</view>
		<view class="main">
			<view class="title">注册</view>
			<view class="slogan">您好，欢迎来到AI聊天室！</view>
			<view class="inputs">
				<view class="inputs-div">
					<input type="text" placeholder="请输入用户名" class="user" placeholder-style="color:#bbb;font-weight:400;" @blur="isUsername"/>
					<view class="wText" v-if="occupyUsername">用户名已被占用</view>
					<view class="wText" v-if="shortUsername">用户名过短</view>
					<image src="../../static/general/right.png" class="ok" v-if="okUsername"></image>
				</view>
				<view class="inputs-div">
					<input type="text" placeholder="请输入邮箱" class="paw" placeholder-style="color:#bbb;font-weight:400;" @blur="isEmail"/>
					<view class="wText" v-if="occupyEmail">邮箱已被占用</view>
					<view class="wText" v-if="invilid">邮箱无效</view>
					<image src="../../static/general/right.png" class="ok" v-if="okEmail"></image>
				</view>
				<view class="inputs-div">
					<input :type="type" placeholder="请输入密码" class="paw" placeholder-style="color:#bbb;font-weight:400;" @blur="isUserPassword"/>
					<view class="wTextPassword" v-if="shortPassword">密码过短</view>
					<image src="../../static/general/right.png" class="okPs" v-if="okPa"></image>
					<image src="../../static/general/look.png" class="look" v-if="look" @tap="lookPassword"></image>
					<image src="../../static/general/unlook.png" class="look" v-if="!look" @tap="lookPassword"></image>
				</view>
				<view class="inputs-div">
					<input :type="type" placeholder="请重新输入密码" class="paw" placeholder-style="color:#bbb;font-weight:400;" @blur="isUserPasswordRe"/>
					<view class="wTextPassword" v-if="shortRe">密码错误</view>
					<image src="../../static/general/right.png" class="okPs" v-if="okRe"></image>
					<image src="../../static/general/look.png" class="look" v-if="look" @tap="lookPassword"></image>
					<image src="../../static/general/unlook.png" class="look" v-if="!look" @tap="lookPassword"></image>
				</view>
			</view>
		</view>
		<view class="submit" @tap="subInfo">注册</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				type:'password',
				occupyUsername: false,
				okUsername: false,
				occupyEmail: false,
				invilid: false,
				okEmail: false,
				shortUsername: false,
				shortPassword: false,
				look: false,
				shortRe: false,
				okRe: false,
				okPa: false,
				userValue: '',
				emailValue: '',
				userPassword: '',
				userPasswordRe: '',
			}
		},
		methods: {
			toSignIn: function(){
				uni.navigateTo({
					url: '../signin/signin',
				});
			},
			lookPassword: function(){
				if(this.look){
					this.type="password";
				}else{
					this.type="text";
				}
				this.look=!this.look;
			},
			isUsername: function(e){
				if(e.detail.cursor>0){
					this.userValue=e.detail.value;
					if(e.detail.cursor>=6){
						this.okUsername=true;
						this.shortUsername=false;
					}else{
						this.okUsername=false;
						this.shortUsername=true;
					}
				}
			},
			isEmail: function(e){
				let reg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/;
				if(e.detail.cursor>0){
					this.emailValue=e.detail.value;
					if(reg.test(this.emailValue)){
						this.okEmail=true;
						this.invilid=false;
					}else{
						this.okEmail=false;
						this.invilid=true;
					}
				}
			},
			isUserPassword: function(e){
				if(e.detail.cursor>0){
					this.userPassword=e.detail.value;
					if(e.detail.cursor>=6){
						this.okPa=true;
						this.shortPassword=false;
						if(this.userPasswordRe && this.userPasswordRe!=this.userPassword){
							this.okRe=false;
							this.shortRe=true;
						}
					}else{
						this.okPa=false;
						this.shortPassword=true;
					}
				}
			},
			isUserPasswordRe: function(e){
				if(e.detail.cursor>0){
					this.userPasswordRe=e.detail.value;
					if(e.detail.cursor>=6 && this.userPasswordRe==this.userPassword){
						this.okRe=true;
						this.shortRe=false;
					}else{
						this.okRe=false;
						this.shortRe=true;
					}
				}
			},
			subInfo: function(){
				if(this.okUsername && this.okEmail && this.okPa && this.okRe){
					console.log(this.userValue);
					console.log(this.emailValue);
					console.log(this.userPassword);
				}
			},
		}
	}
</script>

<style lang="scss">
	@import "../../commons/css/mycss.scss";
	.top-bar{
		border-bottom: 1px solid $uni-border-color;
		.top-bar-right{
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
		padding:54rpx $uni-spacing-row-lg 54rpx;
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
	.inputs-div{
		position: relative;
		.wText,.wTextPassword{
			position: absolute;
			right: 0;
			top: 70rpx;
			font-size: $uni-font-size-base;
			font-weight: 500;
			color: $uni-color-warning;
			float: right;
		}
		image{
			position: absolute;
			right: 0;
			top: 70rpx;
			float: right;
			width: 40rpx;
			height: 40rpx;
		}
		.okPs{
			right: 40rpx;
		}
		.wTextPassword{
			right: 45rpx;
		}
		
	}
</style>
