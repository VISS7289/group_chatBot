export default {
	// 转换时间
	refersh(myurl, atoken, rtoken) {
		return new Promise((resolve, reject) => {
			uni.request({
				url: myurl + '/refersh', // 请求的URL地址
				method: 'POST', // 请求方法为POST
				data: {
					'a_token': atoken, // 请求参数中的a_token
					'r_token': rtoken // 请求参数中的r_token
				},
				success: data => {
					console.log(data.data) // 打印返回的数据
					if (data.data.Code == 1000) { // 判断返回数据的Code字段是否为1000
						try {
							uni.setStorageSync('atoken', data.data.Data[
								0]) // 将返回数据中的第一个元素存储在本地缓存中的'atoken'键中
							uni.setStorageSync('rtoken', data.data.Data[
								1]) // 将返回数据中的第二个元素存储在本地缓存中的'rtoken'键中
						} catch (e) {
							console.log('数据存储出错') // 捕获异常并打印错误信息
							console.log(e)
							uni.navigateTo({
								url: '../signin/signin',
							}) // 导航到 '../signin/signin' 页面
						}
					} else {
						uni.navigateTo({
							url: '../signin/signin',
						}) // 导航到 '../signin/signin' 页面
					}
					resolve(data.data.Code) // 将返回数据的Code字段传递给Promise的resolve函数
				}
			})
		})
	},
}
