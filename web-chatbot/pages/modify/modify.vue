<template>
	<view class="content">
		<view class="top-bar">
			<view class="top-bar-left" @tap="backOne">
				<image src="../../static/general/return.png"></image>
			</view>
			<view class="top-bar-center">
				<view class="title">{{infoType[modifyInfo.type]}}</view>
			</view>
			<view class="top-bar-right" @tap="infoChange">
				<view class="textB">确认</view>
			</view>
		</view>
		<view class="main">
			<view class="modifyInp" v-if="modifyInfo.type==0">
				<view class="inputs-div">
					<input type="text" placeholder="请输入用户名" class="user" placeholder-style="color:#bbb;font-weight:400;"
						@blur="isUsername" />
					<view class="wText" v-if="occupy">用户名已被占用</view>
					<view class="wText" v-if="short">用户名过短</view>
					<image src="../../static/general/right.png" class="ok" v-if="ok"></image>
				</view>
			</view>
			<view class="modifyInp" v-if="modifyInfo.type==1">
				<view class="inputs-div">
					<input :type="type" placeholder="请输入密码" class="paw" placeholder-style="color:#bbb;font-weight:400;"
						@blur="isUserPassword" />
					<view class="wTextPassword" v-if="shortPassword">密码过短</view>
					<image src="../../static/general/right.png" class="okPs" v-if="okPa"></image>
					<image src="../../static/general/look.png" class="look" v-if="look" @tap="lookPassword"></image>
					<image src="../../static/general/unlook.png" class="look" v-if="!look" @tap="lookPassword"></image>
				</view>
				<view class="inputs-div">
					<input :type="type" placeholder="请重新输入密码" class="paw"
						placeholder-style="color:#bbb;font-weight:400;" @blur="isUserPasswordRe" />
					<view class="wTextPassword" v-if="shortRe">密码错误</view>
					<image src="../../static/general/right.png" class="okPs" v-if="okRe"></image>
					<image src="../../static/general/look.png" class="look" v-if="look" @tap="lookPassword"></image>
					<image src="../../static/general/unlook.png" class="look" v-if="!look" @tap="lookPassword"></image>
				</view>
				<view class="vertify-div" v-if="!verifi">
					<input type="text" placeholder="请重新输入验证码" class="paw"
						placeholder-style="color:#bbb;font-weight:400;" @blur="inputVerifi" />
					<view class="vertifySend" @tap="sendVertify">
						<view v-if="send">发送验证码</view>
						<view v-if="unSend" class="unSelect">{{this.time}}s</view>
						<view v-if="reSend">重新发送</view>
					</view>
				</view>
				<view class="tips" v-if="wrong">{{this.errInfo}}</view>
			</view>
			<view class="modifyText" v-if="modifyInfo.type==2">
				<textarea v-model="modifyInfo.userinfo.intr" :placeholder="userRequest" class="modifyContent"
					@blur="isUserIntr"></textarea>
			</view>
			<view class="modifyInp" v-if="modifyInfo.type==3">
				<view class="vertify-div" v-if="!verifi">
					<input type="text" placeholder="请重新输入验证码" class="paw"
						placeholder-style="color:#bbb;font-weight:400;" @blur="inputVerifi" />
					<view class="vertifySend" @tap="sendVertify">
						<view v-if="send">发送验证码</view>
						<view v-if="unSend" class="unSelect">{{this.time}}s</view>
						<view v-if="reSend">重新发送</view>
					</view>
				</view>
				<view class="submit" @tap="subInfo" v-if="!verifi">确认验证码</view>
				<view class="inputs-div" v-if="verifi">
					<input type="text" placeholder="请输入邮箱" class="paw" placeholder-style="color:#bbb;font-weight:400;"
						@blur="isEmail" />
					<view class="wText" v-if="occupyEmail">邮箱已被占用</view>
					<view class="wText" v-if="invilid">邮箱无效</view>
					<image src="../../static/general/right.png" class="ok" v-if="ok"></image>
				</view>
				<view class="vertify-div" v-if="verifi">
					<input type="text" placeholder="请重新输入验证码" class="paw"
						placeholder-style="color:#bbb;font-weight:400;" @blur="inputVerifi2" />
					<view class="vertifySend" @tap="sendVertify2">
						<view v-if="send2">发送验证码</view>
						<view v-if="unSend2" class="unSelect">{{this.time2}}s</view>
						<view v-if="reSend2">重新发送</view>
					</view>
				</view>
				<view class="tips" v-if="wrong">{{this.errInfo}}</view>
			</view>
			<view class="modifyInp" v-if="modifyInfo.type==4">
				<view class="inputs-div">
					<input type="text" placeholder="请输入电话号码" class="paw" placeholder-style="color:#bbb;font-weight:400;"
						@blur="isPhone" />
					<view class="wText" v-if="invilid">电话号码无效</view>
					<image src="../../static/general/right.png" class="ok" v-if="ok"></image>
				</view>
			</view>
			<view class="modifyText" v-if="modifyInfo.type==5">
				<textarea v-model="modifyInfo.userinfo.nick" :placeholder="userRequest" class="modifyContent"
					@blur="isUserNick"></textarea>
			</view>
			<view class="modifyText" v-if="modifyInfo.type==6">
				<textarea v-model="modifyInfo.userRequest" class="modifyContent" :placeholder="userRequest"
					@blur="isAddFriend"></textarea>
			</view>
		</view>
	</view>
</template>

<script>
	import refersh from '../../commons/js/refershToken.js'
	import config from '../../commons/js/config.js'
	export default {
		data() {
			return {
				// infoType 变量用于存储信息类型的数组
				infoType: ['昵称', '密码', '签名', '邮箱', '电话', '备注', '添加好友'],

				// modifyInfo 变量是一个对象，包含两个属性
				modifyInfo: {
					type: 0, // type 用于指示要修改的信息类型，0表示文本，1表示密码，2表示邮箱
					userinfo: {} // userinfo 用于存储用户信息的对象
				},

				// type 变量用于存储当前信息类型的字符串
				type: 'password',

				// occupyUsername 变量用于指示用户名是否已被占用
				occupyUsername: false,

				// okUsername 变量用于指示用户名是否有效
				okUsername: false,

				// occupyEmail 变量用于指示邮箱是否已被占用
				occupyEmail: false,

				// invilid 变量用于指示输入是否无效
				invilid: false,

				// okEmail 变量用于指示邮箱是否有效
				okEmail: false,

				// ok 变量用于指示操作是否成功
				ok: false,

				// look 变量用于指示是否查看详情
				look: false,

				// short 变量用于指示密码是否过短
				short: false,

				// occupy 变量用于指示用户名是否已被占用
				occupy: false,

				// okRe 变量用于指示密码是否有效
				okRe: false,

				// okPa 变量用于指示密码是否有效
				okPa: false,

				// userValue 变量用于存储用户输入的用户名
				userValue: '',

				// emailValue 变量用于存储用户输入的邮箱
				emailValue: '',

				// phoneValue 变量用于存储用户输入的电话号码
				phoneValue: '',

				// userPassword 变量用于存储用户输入的密码
				userPassword: '',

				// userPasswordRe 变量用于存储用户再次输入的密码
				userPasswordRe: '',

				// userRequest 变量用于存储用户的请求
				userRequest: '',

				// atoken 变量用于存储访问令牌
				atoken: '',

				// rtoken 变量用于存储刷新令牌
				rtoken: '',

				// typeRes 变量用于存储类型响应
				typeRes: '',

				// reqRes 变量用于存储请求响应
				reqRes: '',

				// userid 变量用于存储用户ID
				userid: '',

				// verifi 变量用于指示是否已验证
				verifi: false,

				// shortPassword 变量用于指示密码是否过短
				shortPassword: false,

				// shortRe 变量用于指示再次输入的密码是否过短
				shortRe: false,

				// verifiValue 变量用于存储验证码的值
				verifiValue: '',

				// verifiValue2 变量用于存储第二个验证码的值
				verifiValue2: '',

				// time 变量用于存储时间
				time: 0,

				// send 变量用于指示是否发送
				send: true,

				// reSend 变量用于指示是否重新发送
				reSend: false,

				// unSend 变量用于指示是否未发送
				unSend: false,

				// send2 变量用于指示是否发送
				send2: true,

				// reSend2 变量用于指示是否重新发送
				reSend2: false,

				// unSend2 变量用于指示是否未发送
				unSend2: false,

				// okEmail 变量用于指示邮箱是否有效
				okEmail: false,

				// errInfo 变量用于存储错误信息
				errInfo: '',

				// wrong 变量用于指示是否出错
				wrong: false
			}
		},
		// 根据修改信息的种类不同修改页面
		onLoad: function(option) {
			this.modifyInfo = JSON.parse(decodeURIComponent(option.modifyInfo))
			switch (this.modifyInfo.type) {
				case 6:
					this.userRequest = this.modifyInfo.userRequest
					this.ok = true
					break
				case 2:
					this.userRequest = this.modifyInfo.userinfo.intr
					this.ok = true
					break
				case 0:
					this.userRequest = this.modifyInfo.userinfo.name
					break
				case 4:
					this.userRequest = this.modifyInfo.userinfo.phone
					break
				case 5:
					this.userRequest = this.modifyInfo.userinfo.nick
					this.ok = true
					break
				case 3:
					this.userRequest = this.modifyInfo.userinfo.email
					this.ok = false
					break
				case 1:
					this.ok = false
					break
			}
			try {
				this.atoken = uni.getStorageSync('atoken')
				this.rtoken = uni.getStorageSync('rtoken')
				const value = uni.getStorageSync('user')
				this.userid = value.id
			} catch (e) {
				// error
			}
		},
		methods: {
			// 返回上一个页面
			backOne: function() {
				uni.navigateBack({
					delta: 1
				})
			},
			// 输入验证码
			inputVerifi: function(e) {
				this.verifiValue = e.detail.value
			},
			inputVerifi2: function(e) {
				this.verifiValue2 = e.detail.value
			},
			infoChange: function() {
				// 定义一个箭头函数，该函数在后面的代码中被调用
				var fn = () => {
					console.log(2020);
					try {
						let value = uni.getStorageSync('user');
						if (value) {
							value[this.typeRes] = this.reqRes;
							this.modifyInfo[this.typeRes] = this.reqRes;
							uni.setStorageSync('user', value);

							// 获取当前页面和前一个页面
							let pages = getCurrentPages();
							let beforePage = pages[pages.length - 2];

							// 调用前一个页面的 `refersh` 方法，并传递相关参数
							beforePage.$vm.refersh({
								'type': this.typeRes,
								'req': this.reqRes
							});

							// 返回上一个页面
							uni.navigateBack({
								delta: 1
							});
						} else {
							uni.navigateTo({
								url: '../signin/signin'
							});
						}
					} catch (e) {
						console.log(e);
					}
				};

				console.log(this.ok);
				if (this.ok) {
					switch (this.modifyInfo.type) {
						case 6:
							this.addFriend();
							break;
						case 2:
							// 修改签名信息
							this.typeRes = 'intr';
							this.reqRes = this.modifyInfo.userinfo.intr;
							this.update('intr', this.modifyInfo.userinfo.id, this.modifyInfo.userinfo.intr, '', fn);
							break;
						case 0:
							// 修改用户名信息
							this.typeRes = 'name';
							this.reqRes = this.modifyInfo.userinfo.name;
							this.update('username', this.modifyInfo.userinfo.id, this.modifyInfo.userinfo.name, '',
								fn);
							break;
						case 4:
							// 修改电话信息
							this.typeRes = 'phone';
							this.reqRes = this.modifyInfo.userinfo.phone;
							this.update('phone', this.modifyInfo.userinfo.id, this.modifyInfo.userinfo.phone, '', fn);
							break;
						case 5:
							// 修改昵称信息
							this.typeRes = 'nick';
							this.reqRes = this.modifyInfo.userinfo.nick;
							this.updateNick(this.modifyInfo.userinfo.id, this.modifyInfo.userinfo.nick, fn);
							break;
						case 3:
							this.subInfo2();
							break;
						case 1:
							this.subInfoPsw();
							break;
					}
				}
			},

			// 跳转到登录页面
			toSignIn: function() {
				uni.navigateTo({
					url: '../signin/signin',
				});
			},

			// 查看密码
			lookPassword: function() {
				if (this.look) {
					// 如果已经查看密码，则将输入框类型设置为 password
					this.type = 'password';
				} else {
					// 如果未查看密码，则将输入框类型设置为 text
					this.type = 'text';
				}
				// 切换查看密码状态
				this.look = !this.look;
			},

			// 验证用户名是否合法
			isUsername: function(e) {
				// 判断输入框中的字符数是否大于 0
				if (e.detail.cursor > 0) {
					// 更新用户信息中的用户名字段
					this.modifyInfo.userinfo.name = e.detail.value;

					// 判断输入的字符数是否大于等于 6
					if (e.detail.cursor >= 6) {
						this.short = false;

						// 发起网络请求，验证用户名是否已存在
						uni.request({
							url: config.myurl + '/existName?username=' + this.modifyInfo.userinfo.name,
							method: 'GET',
							success: data => {
								console.log(data.data);
								if (data.data.Code == 1002) {
									// 用户名已被占用
									this.occupy = true;
								} else {
									// 用户名未被占用
									this.occupy = false;
								}
								// 根据占用和长度是否合法来判断用户名是否有效
								this.ok = !(this.occupy || this.short);
							},
						});
					} else {
						// 用户名长度不合法
						this.short = true;
						this.occupy = false;
					}
				}
			},

			// 验证邮箱是否合法
			isEmail: function(e) {
				// 正则表达式匹配邮箱格式
				let reg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/;
				if (e.detail.cursor > 0) {
					if (reg.test(e.detail.value)) {
						// 邮箱格式合法
						this.okEmail = true;
						this.modifyInfo.userinfo.email = e.detail.value;
						this.invilid = false;
						this.ok = true;
					} else {
						// 邮箱格式不合法
						this.okEmail = false;
						this.invilid = true;
						this.modifyInfo.userinfo.email = this.userRequest;
						this.ok = false;
					}
				}
			},

			// 验证手机是否合法
			isPhone: function(e) {
				// 正则表达式匹配手机号格式
				let reg = /^([0-9])/;
				if (e.detail.cursor > 0) {
					if (reg.test(e.detail.value) && e.detail.cursor > 6) {
						// 手机号格式合法
						this.ok = true;
						this.modifyInfo.userinfo.phone = e.detail.value;
						this.invilid = false;
					} else {
						// 手机号格式不合法
						this.modifyInfo.userinfo.phone = this.userRequest;
						this.ok = false;
						this.invilid = true;
					}
				}
			},

			// 验证密码是否正确
			isUserPassword: function(e) {
				// 判断输入框中的字符数是否大于 0
				if (e.detail.cursor > 0) {
					this.userPassword = e.detail.value;
					if (e.detail.cursor >= 6) {
						this.okPa = true;
						this.shortPassword = false;
						if (this.userPasswordRe && this.userPasswordRe != this.userPassword) {
							this.okRe = false;
							this.shortRe = true;
						}
					} else {
						this.okPa = false;
						this.shortPassword = true;
					}
				}
			},

			// 密码重复是否正确
			isUserPasswordRe: function(e) {
				// 判断输入框中的字符数是否大于 0
				if (e.detail.cursor > 0) {
					this.userPasswordRe = e.detail.value;
					if (e.detail.cursor >= 6 && this.userPasswordRe == this.userPassword) {
						this.okRe = true;
						this.shortRe = false;
						this.ok = true;
					} else {
						this.okRe = false;
						this.shortRe = true;
						this.ok = false;
					}
				}
			},

			// 用户昵称是否合法
			isUserNick: function(e) {
				// 判断输入框中的字符数是否大于 0
				if (e.detail.cursor > 0) {
					this.modifyInfo.userinfo.nick = e.detail.value;
				} else {
					this.modifyInfo.userinfo.nick = this.userRequest;
				}
			},

			// 添加好友是否合法
			isAddFriend: function(e) {
				this.ok = true;
				if (e.detail.cursor > 0) {
					this.modifyInfo.userRequest = e.detail.value;
				} else {
					this.modifyInfo.userRequest = this.userRequest;
				}
			},

			// 用户信息是否合法
			isUserIntr: function(e) {
				this.ok = true;
				if (e.detail.cursor > 0) {
					this.modifyInfo.userinfo.intr = e.detail.value;
				} else {
					this.modifyInfo.userinfo.intr = this.userRequest;
				}
			},

			// 更新数据
			update(type, optid, data, psw, fn) {
				console.log(fn);
				uni.request({
					url: config.myurl + '/change/update',
					method: 'POST',
					header: {
						'Authorization': 'Bearer ' + this.atoken
					},
					data: {
						'user_id': optid,
						'data': data,
						'type': type,
						'psw': psw
					},
					success: async data => {
						console.log(data.data);
						if (data.data.Code == 1009) {
							// 如果返回的 Code 是 1009，表示令牌过期，需要刷新令牌
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken);
							if (newCode == 1000) {
								// 刷新令牌成功，更新当前令牌和刷新令牌，并重新执行更新操作
								this.atoken = uni.getStorageSync('atoken');
								this.rtoken = uni.getStorageSync('rtoken');
								this.update(type, optid, data, psw);
							} else {
								// 刷新令牌失败，处理错误
								// TODO: 错误处理逻辑
							}
						} else if (data.data.Code == 1000) {
							// 更新成功，执行回调函数
							console.log(1010);
							fn.apply();
						} else {
							// 更新失败，处理错误
							// TODO: 错误处理逻辑
						}
					}
				});
			},

			// 更新昵称
			updateNick(optid, msg, fn) {
				console.log(fn);
				uni.request({
					url: config.myurl + '/change/nick',
					method: 'POST',
					header: {
						'Authorization': 'Bearer ' + this.atoken
					},
					data: {
						'user_id': this.userid,
						'friend_id': optid,
						'msg': msg
					},
					success: async data => {
						console.log(data.data);
						if (data.data.Code == 1009) {
							// 如果返回的 Code 是 1009，表示令牌过期，需要刷新令牌
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken);
							if (newCode == 1000) {
								// 刷新令牌成功，更新当前令牌和刷新令牌，并重新执行更新操作
								this.atoken = uni.getStorageSync('atoken');
								this.rtoken = uni.getStorageSync('rtoken');
								this.updateNick(optid, msg, fn);
							} else {
								// 刷新令牌失败，处理错误
								// TODO: 错误处理逻辑
							}
						} else if (data.data.Code == 1000) {
							// 更新成功，执行回调函数
							console.log(1010);
							fn.apply();
						} else {
							// 更新失败，处理错误
							// TODO: 错误处理逻辑
						}
					}
				});
			},

			// 添加好友
			addFriend: function() {
				uni.request({
					url: config.myurl + '/friend/request',
					method: 'POST',
					header: {
						'Authorization': 'Bearer ' + this.atoken
					},
					data: {
						'user_id': this.modifyInfo.userid,
						'friend_id': this.modifyInfo.friendid,
						'msg': this.modifyInfo.userRequest,
					},
					success: async data => {
						if (data.data.Code == 1009) {
							// 如果返回的 Code 是 1009，表示令牌过期，需要刷新令牌
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken);
							if (newCode == 1000) {
								// 刷新令牌成功，更新当前令牌和刷新令牌，并重新执行添加好友操作
								this.atoken = uni.getStorageSync('atoken');
								this.rtoken = uni.getStorageSync('rtoken');
								this.addFriend();
							} else {
								// 刷新令牌失败，处理错误
								// TODO: 错误处理逻辑
							}
						} else if (data.data.Code == 1000) {
							// 添加好友请求发送成功
							uni.showToast({
								title: '好友请求发送成功',
								icon: 'none',
								duration: 2000
							});

							let pages = getCurrentPages(); // 当前页面
							let beforePage = pages[pages.length - 2];
							uni.navigateBack({
								delta: 1,
								success: function() {
									beforePage.AddSuccess(); // 执行上一页的 AddSuccess 方法
								}
							});
						} else {
							// 添加好友请求发送失败
							uni.showToast({
								title: '好友请求发送失败',
								icon: 'none',
								duration: 2000
							});
						}
					}
				});
			},

			// 上传验证码
			subInfo: function() {
				console.log(this.modifyInfo.userinfo.email);
				console.log(this.verifiValue);
				uni.request({
					url: config.myurl + '/verifiExam',
					method: 'POST',
					data: {
						'email': this.modifyInfo.userinfo.email,
						'verifiCode': this.verifiValue,
					},
					success: data => {
						console.log(data.data);
						if (data.data.Code != 1000) {
							// 验证码上传失败
							this.wrong = true;
							this.errInfo = data.data.Msg;
						} else {
							// 验证码上传成功
							this.wrong = false;
							this.verifi = true;
						}
					}
				});
			},

			// 更改邮箱
			subInfo2: function() {
				console.log(this.verifiValue2);
				uni.request({
					url: config.myurl + '/change/email',
					method: 'POST',
					header: {
						'Authorization': 'Bearer ' + this.atoken
					},
					data: {
						'user_id': this.modifyInfo.userinfo.id,
						'n_email': this.modifyInfo.userinfo.email,
						'verifiCode': this.verifiValue2
					},
					success: async data => {
						console.log(data.data);
						if (data.data.Code == 1009) {
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken);
							if (newCode == 1000) {
								this.atoken = uni.getStorageSync('atoken');
								this.rtoken = uni.getStorageSync('rtoken');
								this.subInfo2();
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
							let value = uni.getStorageSync('user');
							if (value) {
								// 更新本地缓存的用户邮箱
								value.email = this.modifyInfo.userinfo.email;
								this.modifyInfo.email = this.modifyInfo.userinfo.email;
								uni.setStorageSync('user', value);

								let pages = getCurrentPages(); // 当前页面
								let beforePage = pages[pages.length - 2];
								// 调用上一页的刷新方法，更新显示的邮箱
								beforePage.$vm.refersh({
									'type': 'email',
									'req': this.modifyInfo.userinfo.email
								});
								uni.navigateBack({
									delta: 1
								});
							} else {
								uni.navigateTo({
									url: '../signin/signin',
								});
							}
						}
					}
				});
			},

			// 更改密码
			subInfoPsw: function() {
				uni.request({
					url: config.myurl + '/change/psw',
					method: 'POST',
					header: {
						'Authorization': 'Bearer ' + this.atoken
					},
					data: {
						'user_id': this.modifyInfo.userinfo.id,
						'email': this.modifyInfo.userinfo.email,
						'password': this.userPassword,
						're_password': this.userPasswordRe,
						'verifiCode': this.verifiValue
					},
					success: async data => {
						console.log(data.data)
						if (data.data.Code == 1009) {
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								this.atoken = uni.getStorageSync('atoken')
								this.rtoken = uni.getStorageSync('rtoken')
								this.subInfoPsw()
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
							this.wrong = true
							this.errInfo = data.data.Msg
						} else {
							//SUCCESS
							this.wrong = false
							this.verifi = true
							uni.navigateBack({
								delta: 1
							})
						}
					}
				})
			},
			// 发送验证码
			sendVertify: function() {
				if (!this.unSend) {
					this.send = false
					this.reSend = false
					this.unSend = true
					this.time = 60
					var obj = setInterval(() => {
						if (this.time <= 0) {
							this.time = 0
							this.reSend = true
							this.unSend = false
							clearInterval(obj)
						} else {
							this.time = this.time - 1
						}
					}, 1000)
					uni.request({
						url: config.myurl + '/verificationCode',
						method: 'POST',
						data: {
							'to': this.modifyInfo.userinfo.email,
							'username': this.modifyInfo.userinfo.name
						},
						success: data => {
							console.log(data.data)

						}
					})
				} else {
					this.invilid = !this.okEmail
				}
			},
			// 发送验证码
			sendVertify2: function() {
				if (!this.unSend) {
					this.send2 = false
					this.reSend2 = false
					this.unSend2 = true
					this.time2 = 60
					var obj = setInterval(() => {
						if (this.time2 <= 0) {
							this.time2 = 0
							this.reSend2 = true
							this.unSend2 = false
							clearInterval(obj)
						} else {
							this.time2 = this.time2 - 1
						}
					}, 1000)
					uni.request({
						url: config.myurl + '/verificationCode',
						method: 'POST',
						data: {
							'to': this.modifyInfo.userinfo.email,
							'username': this.modifyInfo.userinfo.name
						},
						success: data => {
							console.log(data.data)
						}
					})
				} else {
					this.invilid = !this.okEmail
				}
			},
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

	.tips {
		float: left;
		font-size: $uni-font-size-lg;
		color: $uni-color-warning;
		line-height: 56rpx;
	}

	.main {
		display: flex;
		flex-direction: column;
		align-items: center;
		margin-top: 88rpx;
		padding-top: var(--status-bar-height);
		width: 100%;
		padding: $uni-spacing-col-base $uni-spacing-col-base 0 $uni-spacing-col-base;
		box-sizing: border-box;

		.modifyText {
			width: 100%;
		}

		.modifyContent {
			height: 386rpx;
			width: 100%;
			background: $uni-bg-color-grey;
			border-radius: $uni-border-radius-base;
			font-size: $uni-font-size-lg;
			color: $uni-text-color;
			line-height: 44rpx;
			padding: 16rpx 22rpx;
			box-sizing: border-box;
		}

		.modifyInp {
			width: 100%;
			padding-top: 8rpx;

			input {
				padding-top: 40rpx;
				height: 88rpx;
				font-size: $uni-font-size-lg;
				font-weight: 500;
				color: $uni-text-color;
				line-height: 88rpx;
				border-bottom: 1px solid $uni-border-color;
			}

		}

		.vertify-div {
			display: flex;
			flex-direction: cow;
			align-items: center;
			justify-content: center;
			margin-bottom: 120rpx;

			.vertifySend {
				padding-top: 48rpx;
				padding-left: $uni-spacing-row-base;
				font-weight: 500;
				color: $uni-text-color;
				width: 200rpx;
				text-align: center;
			}

			.unSelect {
				color: $uni-text-color-disable;
			}
		}

		.inputs-div {
			position: relative;

			.wText,
			.wTextPassword {
				position: absolute;
				right: 0;
				top: 70rpx;
				font-size: $uni-font-size-base;
				font-weight: 500;
				color: $uni-color-warning;
				float: right;
			}

			image {
				position: absolute;
				right: 0;
				top: 70rpx;
				float: right;
				width: 40rpx;
				height: 40rpx;
			}

			.okPs {
				right: 40rpx;
			}

			.wTextPassword {
				right: 45rpx;
			}

		}
	}
</style>
