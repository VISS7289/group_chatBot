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
				<view class="vertify-div">
					<input :type="type" placeholder="请重新输入验证码" class="paw" placeholder-style="color:#bbb;font-weight:400;" @blur="inputVerifi"/>
					<view class="vertifySend" @tap="sendVertify">
						<view v-if="send">发送验证码</view>
						<view v-if="unSend" class="unSelect">{{this.time}}s</view>
						<view v-if="reSend">重新发送</view>
					</view>
				</view>
			</view>
			<view class="tips" v-if="wrong">{{this.errInfo}}</view>
		</view>
		<view class="submit" @tap="subInfo">注册</view>
	</view>
</template>

<script>
	import config from '../../commons/js/config.js';
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
				verifiValue: '',
				wrong: false,
				send: true,
				reSend: false,
				unSend: false,
				time: 60,
				errInfo: ''
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
						this.shortUsername=false;
						uni.request({
							url: config.myurl+'/existName?username='+this.userValue,
							method: 'GET',
							success:(data)=> {
								console.log(data.data)
								if(data.data.Code==1002){
									this.occupyUsername=true;
								}else{
									this.occupyUsername=false;
								}
								this.okUsername=!(this.occupyUsername||this.shortUsername);
							},
						})
					}else{
						this.shortUsername=true;
						this.occupyUsername=false;
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
			inputVerifi: function(e){
				this.verifiValue=e.detail.value;
			},
			subInfo: function(){
				this.wrong=!this.wrong;
				if(this.okUsername && this.okEmail && this.okPa && this.okRe){
					uni.request({
						url: config.myurl+'/register',
						method: 'POST',
						data: {
							"username": this.userValue,
							"password": this.userPassword,
							"re_password": this.userPasswordRe,
							"email": this.emailValue,
							"verifiCode": this.verifiValue
						},
						success:(data)=> {
							console.log(data.data)
							if(data.data.Code!=1000){
								this.wrong=true;
								this.errInfo=data.data.Msg;
							}else{
								this.wrong=false;
							}
						}
					})
				}else{
					this.shortUsername=!this.okUsername;
					this.invilid=!this.okEmail;
					this.shortPassword=!this.okPa;
					this.shortRe=!this.okRe;
				}
				
			},
			sendVertify: function(){
				if(this.okEmail&&!this.unSend){
					this.send=false;
					this.reSend=false;
					this.unSend=true;
					var obj = setInterval(() => {
						if(this.time<=0){
							this.time=0;
							this.reSend=true;
							this.unSend=false;
							clearInterval(obj);
						}else{
							this.time=this.time-1;
						}
					}, 1000);
					uni.request({
						url: config.myurl+'/verificationCode',
						method: 'POST',
						data: {
							"to": this.emailValue,
							"username": "duan666"
						},
						success:(data)=> {
							console.log(data.data)
							if(data.data.Code==1000){
								uni.setStorageSync('aToken', data.data.Data[0]);
								uni.setStorageSync('rToken', data.data.Data[1]);
							}
						}
					})
				}else{
					this.invilid=!this.okEmail;
				}
			},
			timeReduce: function(t){
				
				return
			}
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
		padding:54rpx $uni-spacing-row-lg 120rpx;
		height: 850rpx;
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
	.vertify-div{
		display: flex;
		flex-direction: cow;
		align-items: center;
		.vertifySend{
			padding-top: 48rpx;
			padding-left: $uni-spacing-row-base;
			font-weight: 500;
			color: $uni-text-color;
			width: 200rpx;
			text-align: center;
		}
		.unSelect{
			color: $uni-text-color-disable;
		}
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
