<script>
// stream
export default {
	data: function () {
		return {
			photos: [],
			displayUsername: '',
            token: sessionStorage.getItem('token'),
		}
	},
	methods: {
		async getMyStream() {
			try {
                let response = await this.$axios.get('/users/' + this.token + "/stream", {headers: {'Authorization': `${sessionStorage.getItem('token')}`}});
				this.photos = response.data === null ? [] : response.data;
				console.log(this.photos)
				this.displayUsername = sessionStorage.getItem('username');
			} catch (error) {
				const status = error.response.status;
        		const reason = error.response.data;
                this.errormsg = `Status ${status}: ${reason}`;
            }
		},
	},
	async mounted() {
		await this.getMyStream()
	}
}
</script>

<template>
	<div class="container-fluid">
		<div class="row">
			<div class="col">
				<div class="stream-header">
				This is {{ displayUsername }}'s stream
				</div>
			</div>
		</div>
		<div class="row">
			<Photo v-for="photo in photos"
			:key="photo.Pid"
			:pid="photo.Pid"
			:user_id="photo.user_id"
			:username="photo.username"
			:date="photo.date"
			:likesListParent="photo.likes"
			:commentsListParent="photo.comments"
			:isOwner="false"
			/>
		</div>
		<div v-if="photos.length === 0" class="row">
			<h1 class="d-flex justify-content-center mt-5" style="color: rgb(0, 0, 0);">There's no content yet, follow somebody!</h1>
		</div>
	</div>
</template>