class TimeAgo extends HTMLElement{
    connectedCallback(){
        const el = this.querySelector("time")
        if(el ===null|| el.dateTime ==="" ){
            return 
        }

        el.textContent = neatime(el.dateTime)
    }
}

customElements.define("time-ago",TimeAgo)

function neatime(text) {
	var now = new Date(),
		date = new Date(text),
		diff = (now.getTime() - date.getTime()) / 1000;
	if (diff <= 60) {
		text = 'Just now';
	}
	else if ( (diff /= 60) < 60 ) {
		text = (diff|0) + 'm ago';
	}
	else if ( (diff /= 60) < 24 ) {
		text = (diff|0) + 'h ago';
	}
	else if ( (diff /= 24) < 7 ) {
		text = (diff|0) + 'd ago';
	}
	else {
		text = String(date).split(' ')[1] + ' ' + date.getDate();
		if (diff>182 && now.getFullYear()!==date.getFullYear()) {
			text += ', ' + date.getFullYear();
		}
	}console.log(text);
	return text;
    
}