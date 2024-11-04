<script>
export default {	
	data(){
		return{
            token: sessionStorage.getItem('token'),
			text:"",
		}
	},

	props:['modalID','comments','isOwner','pid'],

	methods: {

		async commentPhoto() {
            try {
                let response = await this.$axios.post(`/photos/${this.pid}/comments/${this.token}`, this.text, {headers: {'Authorization': `${sessionStorage.getItem('token')}`, 'Content-Type': 'text/plain'}});
                let comment = response.data;
                this.$emit('addComment', comment); 
                console.log("new comment:" ,this.comments);
                this.text = "";
            } catch(error) {
                const status = error.response.status;
                const reason = error.response.data;
                this.errormsg = `Status` + status + `: ` + reason;
                alert(this.errormsg);
            }
        },
        async uncommentPhoto() {
            try {
                await this.$axios.delete(`/photos/${this.pid}/comments/${this.token}/`, {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
                this.$emit('removeComment', comment_id); 
            } catch(error) {
                const status = error.response.status;
                const reason = error.response.data;
                this.errormsg = `Status` + status + `: ` + reason;
                alert(this.errormsg);
            }
        },

	},
    mounted() {

        console.log("Comments List:", this.comments);
    }
}
</script>
