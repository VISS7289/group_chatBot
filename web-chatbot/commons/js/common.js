export default{
	refersh: function(theurl){
		uni.request({
			url: theurl+'/refersh',
			method: 'POST',
			header: { 'Authorization': 'Bearer ' + uni.getStorageSync('rToken') },
			data: {
				"a_token": uni.getStorageSync('aToken'),
				"r_token": uni.getStorageSync('rToken'),
			},
			success:(data)=> {
				console.log(data);
				uni.setStorageSync('aToken', data.data.Data[0]);
				uni.setStorageSync('rToken', data.data.Data[1]);
			}
		})
	}
}