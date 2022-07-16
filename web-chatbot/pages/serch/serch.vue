<template>
	<view class="content">
		<view class="top-bar">
			<view class="serch-div">
				<input type="serch" placeholder="搜索用户/群" class="serch" placeholder-style="color:#bbb;font-weight:400;" @input="searchKey"/>
				<image src="../../static/index/search.png" class="serch-img"></image>
			</view>
			<view class="top-bar-right" @tap="toSignUp">
				<view class="text">取消</view>
			</view>
		</view>
		<view class="main">
			<view class="serchUser">
				<view class="titleUser">用户</view>
				<view class="listUser" v-for="(item,index) in userArr" :key="index">
					<image :src="item.imgurl"></image>
					<view class="nameUser" v-html="item.name"></view>
					<view class="rightBt info" v-if="dis">发消息</view>
					<view class="rightBt friend">加好友</view>
				</view>
			</view>
			<view class="serchUser">
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
			</view>
		</view>
	</view>
</template>

<script>
	import datas from '../../commons/js/datas.js';
	export default {
		data() {
			return {
				dis: false,
				userArr: [],
			}
		},
		methods: {
			searchKey: function(e){
				this.userArr=[];
				let serchVal=e.detail.value;
				if(serchVal.length>0){
					this.userDel(serchVal);
				}
				
			},
			userDel: function(e){
				let myArr=datas.friends();
				let light=eval("/"+e+"/g");
				for(let i=0;i<myArr.length;i++){
					if(myArr[i].name.search(e)!=-1){
						this.friendDel(myArr[i]);
						console.log(myArr[i].tip);
						myArr[i].imgurl='../../static/index/'+myArr[i].imgurl;
						myArr[i].name=myArr[i].name.replace(light,"<span style='color:#4AAAFF;'>"+e+"</span>");
						this.userArr.push(myArr[i]);
					}
				}
			},
			friendDel: function(e){
				let friendArr=datas.friendTable();
				for(let i=0;i<friendArr.length;i++){
					if(friendArr[i].friendid==e.id){
						e.tip=1;
					}else{
						e.tip=0;
					}
				}
			}
		}
	}
</script>

<style lang="scss">
	@import "../../commons/css/mycss.scss";
	.top-bar{
		background-color: rgba(255,255,255,0.98);
		border-bottom: 1px solid $uni-border-color;
		.serch-div{
			position: absolute;
			top: 0;
			left: 0;
			width: 100%;
			z-index: -1;
			padding: 14rpx;
			.serch{
				width: 540rpx;
				height: 60rpx;
				background-color: $uni-bg-color-grey;
				border-radius: 10rpx;
				padding: 0rpx 50rpx 0rpx 18rpx;
			}
			.serch-img{
				position: absolute;
				right: 165rpx;
				top: 25rpx;
				width: 40rpx;
				height: 40rpx;
			}
		}
		.top-bar-right{
			padding-right: 0rpx;
			.text{
				font-size: 40rpx;
				font-weight: 500;
				color: $uni-text-color;
				line-height: 80rpx;
			}
		}
	}
	.main{
		width: 100%;
		padding-top: 88rpx;
		padding-left: 32rpx;
		.serchUser{
			width: 100%;
			padding-top: $uni-spacing-col-base;
			.titleUser{
				font-size: 44rpx;
				font-weight: 600;
				color: $uni-text-color;
				line-height: 60rpx;
			}
			.listUser{
				display: flex;
				flex-direction: row;
				align-items: center;   
				//justify-content: center;
				width: 100%;
				margin-top: 40rpx;
				image{
					width: 80rpx;
					height: 80rpx;
					float: left;
					border-radius: $uni-border-radius-base;
				}
				.nameUser{
					width: 500rpx;
					padding-left: 32rpx;
					font-size: 36rpx;
					font-weight: 400rpx;
					line-height: 50rpx;
					color: $uni-text-color;
				}
				.rightBt{
					width: 120rpx;
					height: 40rpx;
					border-radius: $uni-border-radius-base;
					
					font-size: $uni-font-size-sm;
					line-height: 40rpx;
					text-align: center;
				}
				.info{
					background-color: $uni-color-primary;
					color: $uni-text-color;
				}
				.friend{
					background-color: $uni-bg-color-grey;
					color: $uni-color-success;
				}
			}
		}
	}
</style>
