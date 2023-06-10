export default {
	//转换时间
	dateTime: function(d) {
		let old = new Date(d); // 创建一个Date对象，表示传入的日期
		let now = new Date(); // 创建一个Date对象，表示当前日期
		let delta = (now.getFullYear() - old.getFullYear()) * 365 +
			(now.getMonth() - old.getMonth()) * 30 +
			(now.getDate() - old.getDate()); // 计算当前日期与传入日期之间的天数差

		if (delta < 1) {
			if (old.getMinutes() < 10) {
				return old.getHours() + ':0' + old.getMinutes(); // 返回小时和分钟，格式为hh:mm
			}
			return old.getHours() + ':' + old.getMinutes(); // 返回小时和分钟，格式为hh:mm
		} else if (delta < 2) {
			return '昨天'; // 返回字符串"昨天"
		} else if (delta < 3) {
			return '前天'; // 返回字符串"前天"
		} else if (delta < 365) {
			return old.getFullYear() + '/' + old.getMonth() + '/' + old.getDate(); // 返回年月日，格式为yyyy/mm/dd
		}

		return '超过一年'; // 返回字符串"超过一年"
	},

	//转换时间
	dateTime2: function(d) {
		let old = [
			parseInt(d.slice(0, 4)), // 年份
			parseInt(d.slice(5, 7)), // 月份
			parseInt(d.slice(8, 10)), // 日期
			parseInt(d.slice(11, 13)), // 小时
			parseInt(d.slice(14, 16)), // 分钟
			parseInt(d.slice(17, 19)) // 秒钟
		];

		let nowTime = new Date();
		let now = [
			nowTime.getFullYear(), // 当前年份
			nowTime.getMonth() + 1, // 当前月份
			nowTime.getDate(), // 当前日期
			nowTime.getHours(), // 当前小时
			nowTime.getMinutes(), // 当前分钟
			nowTime.getSeconds() // 当前秒钟
		];

		let delta = (now[0] - old[0]) * 365 + (now[1] - old[1]) * 30 + (now[2] - old[2]); // 计算当前日期与传入日期之间的天数差

		if (delta < 1) {
			if (old[4] < 10) {
				return old[3] + ':0' + old[4]; // 返回小时和分钟，格式为hh:mm
			}
			return old[3] + ':' + old[4]; // 返回小时和分钟，格式为hh:mm
		} else if (delta < 2) {
			if (old[4] < 10) {
				return '昨天' + old[3] + ':0' + old[4]; // 返回"昨天hh:mm"
			}
			return '昨天' + old[3] + ':' + old[4]; // 返回"昨天hh:mm"
		} else if (delta < 3) {
			return '前天'; // 返回字符串"前天"
		} else if (delta < 365) {
			return old[0] + '/' + old[1] + '/' + old[2]; // 返回年月日，格式为yyyy/mm/dd
		}

		return '超过一年'; // 返回字符串"超过一年"
	},

	//2022-10-03T14:30:55Z padStart
	getNewTime: function() {
		let nowTime = new Date(); // 创建一个Date对象，表示当前日期和时间
		let m = nowTime.getMonth() + 1; // 获取当前月份（月份范围是0-11，所以需要加1）

		// 构建时间字符串，格式为YYYY-MM-DDTHH:MM:SSZ
		return nowTime.getFullYear() + '-' + m.toString().padStart(2, '0') + '-' + nowTime.getDate().toString()
			.padStart(2, '0') + 'T' +
			nowTime.getHours().toString().padStart(2, '0') + ':' + nowTime.getMinutes().toString().padStart(2,
				'0') + ':' + nowTime.getMonth().toString().padStart(2, '0') + 'Z';
	},

	//聊天时间
	chatTime: function(d) {
		let old = new Date(d); // 创建一个Date对象，表示传入的日期
		let now = new Date(); // 创建一个Date对象，表示当前日期
		let delta = (now.getFullYear() - old.getFullYear()) * 365 +
			(now.getMonth() - old.getMonth()) * 30 +
			(now.getDate() - old.getDate()); // 计算当前日期与传入日期之间的天数差

		if (delta < 1) {
			if (old.getMinutes() < 10) {
				return old.getHours() + ':0' + old.getMinutes(); // 返回小时和分钟，格式为hh:mm
			}
			return old.getHours() + ':' + old.getMinutes(); // 返回小时和分钟，格式为hh:mm
		} else if (delta < 2) {
			if (old.getMinutes() < 10) {
				return '昨天 ' + old.getHours() + ':0' + old.getMinutes(); // 返回"昨天 hh:mm"
			}
			return '昨天 ' + old.getHours() + ':' + old.getMinutes(); // 返回"昨天 hh:mm"
		} else if (delta < 3) {
			return '前天' + old.getHours() + ':0' + old.getMinutes(); // 返回"前天 hh:mm"
		} else if (delta < 365) {
			return old.getFullYear() + '/' + old.getMonth() + '/' + old.getDate(); // 返回年月日，格式为yyyy/mm/dd
		}

		return '超过一年'; // 返回字符串"超过一年"
	},
	//间隔时间
	spaceTime: function(oldTime, nowTime) {
		let old = new Date(oldTime); // 创建一个Date对象，表示传入的旧时间
		let now = new Date(nowTime); // 创建一个Date对象，表示传入的当前时间
		let delta = (now.getFullYear() - old.getFullYear()) * 60 * 24 * 30 * 365 +
			(now.getMonth() - old.getMonth()) * 60 * 24 * 30 +
			(now.getDate() - old.getDate()) * 60 * 24 +
			(now.getHours() - old.getHours()) * 60 +
			(now.getMinutes() - old.getMinutes()); // 计算当前时间与旧时间之间的时间差，单位为分钟

		if (delta > 5) {
			return oldTime; // 如果时间差大于5分钟，则返回旧时间
		}
		return ''; // 如果时间差小于等于5分钟，则返回空字符串
	},
	//间隔时间
	spaceTime2: function(time1, time2) {
		let time1d = [parseInt(time1.slice(0, 4)), parseInt(time1.slice(5, 7)), parseInt(time1.slice(8, 10)),
			parseInt(time1.slice(11, 13)), parseInt(time1.slice(14, 16)), parseInt(time1.slice(17, 19))
		]; // 解析时间字符串 time1，将年、月、日、小时、分钟、秒钟分别存储到数组 time1d 中
		let time2d = [parseInt(time2.slice(0, 4)), parseInt(time2.slice(5, 7)), parseInt(time2.slice(8, 10)),
			parseInt(time2.slice(11, 13)), parseInt(time2.slice(14, 16)), parseInt(time2.slice(17, 19))
		]; // 解析时间字符串 time2，将年、月、日、小时、分钟、秒钟分别存储到数组 time2d 中
		let delta = (time2d[0] - time1d[0]) * 60 * 24 * 30 * 365 +
			(time2d[1] - time1d[1]) * 60 * 24 * 30 +
			(time2d[2] - time1d[2]) * 60 * 24 +
			(time2d[3] - time1d[3]) * 60 +
			(time2d[4] - time1d[4]); // 计算 time1 和 time2 之间的时间差，以分钟为单位
		return delta; // 返回时间差
	},
	//比较时间 2022-10-03T14:30:55Z
	compareTime: function(time1, time2) {
		let time1d = [parseInt(time1.slice(0, 4)), parseInt(time1.slice(5, 7)), parseInt(time1.slice(8, 10)),
			parseInt(time1.slice(11, 13)), parseInt(time1.slice(14, 16)), parseInt(time1.slice(17, 19))
		]; // 解析时间字符串 time1，将年、月、日、小时、分钟、秒钟分别存储到数组 time1d 中
		let time2d = [parseInt(time2.slice(0, 4)), parseInt(time2.slice(5, 7)), parseInt(time2.slice(8, 10)),
			parseInt(time2.slice(11, 13)), parseInt(time2.slice(14, 16)), parseInt(time2.slice(17, 19))
		]; // 解析时间字符串 time2，将年、月、日、小时、分钟、秒钟分别存储到数组 time2d 中

		for (let i = 0; i < 6; i++) {
			if (time1d[i] > time2d[i]) {
				return 1; // 如果 time1 在第 i 个时间单位上大于 time2，则返回 1
			}
			if (time1d[i] < time2d[i]) {
				return 2; // 如果 time1 在第 i 个时间单位上小于 time2，则返回 2
			}
		}
		return 0; // 如果 time1 和 time2 完全相等，则返回 0
	},
	mySortByTime: function(arr, obj, type) {
		//type 1 降序 0 升序
		// console.log(arr[0][obj])
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
