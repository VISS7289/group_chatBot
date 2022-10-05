export default{
	//转换时间
	refersh(myurl,atoken,rtoken){
		return new Promise((resolve, reject)=>{
			uni.request({
				url: myurl + '/refersh',
				method: 'POST',
				data: {
					'a_token': atoken,
					'r_token': rtoken
				},
				success: data => {
					console.log(data.data)
					if (data.data.Code == 1000) {
						try {
							uni.setStorageSync('atoken', data.data.Data[0])
							uni.setStorageSync('rtoken', data.data.Data[1])
						} catch (e) {
							console.log('数据存储出错')
							console.log(e)
							uni.navigateTo({ url: '../signin/signin', })
						}
					} else {
						uni.navigateTo({ url: '../signin/signin', })
					}
					resolve(data.data.Code)
				}
			})
		})
	},
}