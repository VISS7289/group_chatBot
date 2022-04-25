export default{
	//转换时间
	dateTime: function(d) {
		let old = new Date(d);
		let now = new Date();
		let delta = (now.getFullYear() - old.getFullYear())*365
		+ (now.getMonth() - old.getMonth())*30
		+ (now.getDate() - old.getDate());
		if(delta < 1) {
			if(old.getMinutes()<10) {
				return old.getHours()+':0'+old.getMinutes();
			}
			return old.getHours()+':'+old.getMinutes();
		}
		else if(delta < 2) {
			return '昨天';
		}
		else if(delta < 3) {
			return '前天';
		}
		else if(delta < 365) {
			return old.getFullYear()+'/'+old.getMonth()+'/'+old.getDate();
		}
		return '超过一年';
	},
	//聊天时间
	chatTime: function(d) {
		let old = new Date(d);
		let now = new Date();
		let delta = (now.getFullYear() - old.getFullYear())*365
		+ (now.getMonth() - old.getMonth())*30
		+ (now.getDate() - old.getDate());
		if(delta < 1) {
			if(old.getMinutes()<10) {
				return old.getHours()+':0'+old.getMinutes();
			}
			return old.getHours()+':'+old.getMinutes();
		}
		else if(delta < 2) {
			if(old.getMinutes()<10) {
				return '昨天 '+old.getHours()+':0'+old.getMinutes();
			}
			return '昨天 '+old.getHours()+':'+old.getMinutes();
		}
		else if(delta < 3) {
			return '前天'+old.getHours()+':0'+old.getMinutes();
		}
		else if(delta < 365) {
			return old.getFullYear()+'/'+old.getMonth()+'/'+old.getDate();
		}
		return '超过一年';
	},
	//间隔时间
	spaceTime: function(oldTime,nowTime) {
		let old = new Date(oldTime);
		let now = new Date(nowTime);
		let delta = (now.getFullYear() - old.getFullYear())*60*24*30*365
		+ (now.getMonth() - old.getMonth())*60*24*30
		+ (now.getDate() - old.getDate())*60*24
		+ (now.getHours() - old.getHours())*60
		+ (now.getMinutes() - old.getMinutes());
		if(delta > 5) {
			return oldTime;
		}
		return '';
	}
}