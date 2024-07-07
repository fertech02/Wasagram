<script>
const token = sessionStorage.getItem('authToken');
export default {

    data() {
        return {
            photo: null,
            errormsg: null,
            detailedmsg: null,
            successmsg: null,
        }
    },

    methods:
    {
        async uploadFile() {
			this.images = this.$refs.file.files[0]
		},
		async submitFile() {
			if (this.photo === null) {
				this.errormsg = "Please select a file to upload."
			} else {
				try {
					const response = await this.$axios.post("/photos/", {
                        photo: this.photo,
                    }, {
                        headers: {
                            "Authorization": `Bearer ${token}`
                        }
                    })
					this.successmsg = "Photo uploaded successfully."
				} catch (e) {
					if (e.response && e.response.status === 400) {
						this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
						this.detailedmsg = null;
					} else if (e.response && e.response.status === 500) {
						this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
						this.detailedmsg = e.toString();
					} else {
						this.errormsg = e.toString();
						this.detailedmsg = null;
					}
				}
			}
		},
    }
}
</script>   

<template>
    <div>
        <h1>Upload Photo</h1>
        <div>
            <input type="file" ref="file" v-on:change="uploadFile" />
            <button v-on:click="submitFile">Upload</button>
        </div>
        <div v-if="errormsg">
            <p>{{ errormsg }}</p>
            <p>{{ detailedmsg }}</p>
        </div>
        <div v-if="successmsg">
            <p>{{ successmsg }}</p>
        </div>
    </div>
</template>
  