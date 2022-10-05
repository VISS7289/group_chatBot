export default {
	//转换时间
	dateTime: function(d) {
		let old = new Date(d)
		let now = new Date()
		let delta = (now.getFullYear() - old.getFullYear()) * 365 +
			(now.getMonth() - old.getMonth()) * 30 +
			(now.getDate() - old.getDate())
		if (delta < 1) {
			if (old.getMinutes() < 10) {
				return old.getHours() + ':0' + old.getMinutes()
			}
			return old.getHours() + ':' + old.getMinutes()
		} else if (delta < 2) {
			return '昨天'
		} else if (delta < 3) {
			return '前天'
		} else if (delta < 365) {
			return old.getFullYear() + '/' + old.getMonth() + '/' + old.getDate()
		}
		return '超过一年'
	},
	//转换时间
	dateTime2: function(d) {
		let old = [parseInt(d.slice(0, 4)), parseInt(d.slice(5, 7)), parseInt(d.slice(8, 10)),
			parseInt(d.slice(11, 13)), parseInt(d.slice(14, 16)), parseInt(d.slice(17, 19))
		]
		let nowTime = new Date()
		let now = [nowTime.getFullYear(), nowTime.getMonth(), nowTime.getDate(), nowTime.getHours(),
			nowTime.getMinutes(), nowTime.getMonth()
		]
		let delta = (now[0] - old[0]) * 365 + (now[1] - old[1]) * 30 + (now[2] - old[2])
		if (delta < 1) {
			if (old[4] < 10) {
				return old[3] + ':0' + old[4]
			}
			return old[3] + ':' + old[4]
		} else if (delta < 2) {
			return '昨天'
		} else if (delta < 3) {
			return '前天'
		} else if (delta < 365) {
			return old[0] + '/' + old[1] + '/' + old[2]
		}
		return '超过一年'
	},
	//聊天时间
	chatTime: function(d) {
		let old = new Date(d)
		let now = new Date()
		let delta = (now.getFullYear() - old.getFullYear()) * 365 +
			(now.getMonth() - old.getMonth()) * 30 +
			(now.getDate() - old.getDate())
		if (delta < 1) {
			if (old.getMinutes() < 10) {
				return old.getHours() + ':0' + old.getMinutes()
			}
			return old.getHours() + ':' + old.getMinutes()
		} else if (delta < 2) {
			if (old.getMinutes() < 10) {
				return '昨天 ' + old.getHours() + ':0' + old.getMinutes()
			}
			return '昨天 ' + old.getHours() + ':' + old.getMinutes()
		} else if (delta < 3) {
			return '前天' + old.getHours() + ':0' + old.getMinutes()
		} else if (delta < 365) {
			return old.getFullYear() + '/' + old.getMonth() + '/' + old.getDate()
		}
		return '超过一年'
	},
	//间隔时间
	spaceTime: function(oldTime, nowTime) {
		let old = new Date(oldTime)
		let now = new Date(nowTime)
		let delta = (now.getFullYear() - old.getFullYear()) * 60 * 24 * 30 * 365 +
			(now.getMonth() - old.getMonth()) * 60 * 24 * 30 +
			(now.getDate() - old.getDate()) * 60 * 24 +
			(now.getHours() - old.getHours()) * 60 +
			(now.getMinutes() - old.getMinutes())
		if (delta > 5) {
			return oldTime
		}
		return ''
	},
	//比较时间 2022-10-03T14:30:55Z
	compareTime: function(time1, time2) {
		let time1d = [parseInt(time1.slice(0, 4)), parseInt(time1.slice(5, 7)), parseInt(time1.slice(8, 10)),
			parseInt(time1.slice(11, 13)), parseInt(time1.slice(14, 16)), parseInt(time1.slice(17, 19))
		]
		let time2d = [parseInt(time2.slice(0, 4)), parseInt(time2.slice(5, 7)), parseInt(time2.slice(8, 10)),
			parseInt(time2.slice(11, 13)), parseInt(time2.slice(14, 16)), parseInt(time2.slice(17, 19))
		]
		for (let i = 0; i < 6; i++) {
			if (time1d[i] > time2d[i]) {
				return 1
			}
			if (time1d[i] < time2d[i]) {
				return 2
			}
		}
		return 0
	},
	mySortByTime: function(arr, obj, type) {
		//type 1 降序 0 升序
		console.log(arr[0][obj])
		if (type == 0) {
			for (let i = 1; i < arr.length; i++) {
				for (let j = i; j > 0; j--) {
					if (this.compareTime(arr[j][obj], arr[j - 1][obj]) == 1) {
						let s = arr[j][obj]
						arr[j][obj] = arr[j - 1][obj]
						arr[j - 1][obj] = s
					}
				}
			}
		} else {
			for (let i = 1; i < arr.length; i++) {
				for (let j = i; j > 0; j--) {
					if (this.compareTime(arr[j][obj], arr[j - 1][obj]) == 2) {
						let s = arr[j][obj]
						arr[j][obj] = arr[j - 1][obj]
						arr[j - 1][obj] = s
					}
				}
			}
		}
		return arr
	}
}
