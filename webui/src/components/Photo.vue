<script>
export default {
	data() {
		return {
            errormsg: null,
			photoURL: "",
			liked: false,
            likes: [],
            comments: [],
		}
	},

	props: ['pid','ownerID','username','date','likesListParent','commentsListParent','isOwner'], 

	methods: {
		getPhoto() {
            this.photoURL = API_URL + '/photos/' + this.pid;
		},

		async deletePhoto() {
            try {
                await this.$axios.delete(`/photos/${this.pid}`);
                this.$emit("removePhoto", this.pid);

            } catch (error) {
                const status = error.response.status;
                const reason = error.response.data;
                this.errormsg = `Status ${status}: ${reason}`;
            }
		},

		visitAuthorProfile() {
            this.$router.push('/users' + this.ownerID + '/profile');
		},

		async likeToggle() {
            try {
                if (!this.liked) {
                    await this.$axios.delete(`/photos/${this.pid}/likes/${sessionStorage.getItem('token')}`);
                    this.likes = this.likes.filter(user => user.user_id != sessionStorage.getItem('token'));
                } else {
                    await this.$axios.post(`/photos/${this.pid}/likes/${sessionStorage.getItem('token')}`);
                    this.likes.unshift({user_id: sessionStorage.getItem('token')});
                }
                this.liked = !this.liked;
            } catch (error) {
                const status = error.response.status;
                const reason = error.response.data;
                this.errormsg = `Status ${status}: ${reason}`;
            }
    	},
        

		removeCommentFromList(comment_id) {
			this.comments = this.comments.filter(comment => comment.comment_id != comment_id);
		},

		addCommentToList(comment){
			this.comments.unshift(comment); // at the beginning of the list
		},
        
	},
	async mounted() {
        this.getPhoto()
        if (this.likesListParent != null) {
            this.likes = this.likesListParent
        }
        if (this.commentsListParent != null) {
            this.comments = this.commentsListParent
        }
		this.liked = this.likes.some(user => user.user_id == sessionStorage.getItem('token'));
	},
}
</script>

<template>
	<div class="container-fluid mt-3 mb-5 ">

        <UserModal
        :modalID="'likesModal'+pid" 
		:usersList="likes"
        />

        <CommentModal
        :modalID="'commentModal'+pid" 
		:comments="comments" 
		:isOwner="isOwner" 
		:pid="pid"
		@removeComment="removeCommentFromList"
		@addComment="addCommentToList"
		/>

        <div class="d-flex flex-row justify-content-center">
            <div class="card my-card">
                <div class="d-flex justify-content-end">
                    <button v-if="isOwner" class="my-trnsp-btn my-dlt-btn me-2" @click="deletePhoto">
						<!--trash bin-->
						<i class="fa-solid fa-trash w-100 h-100"></i>
					</button>
                </div>
                <!--photo-->
                <div class="d-flex justify-content-center photo-background-color">
                    <img :src="photoURL" class="card-img-top img-fluid">
                </div>
                <div class="card-body">
                    <div class="container">
                        <div class="d-flex flex-row justify-content-end align-items-center mb-2">
                            <!--like-->
                            <button class="my-trnsp-btn m-0 p-1 d-flex justify-content-center align-items-center">
                                <i @click="likeToggle" :class="'me-1 my-heart-color w-100 h-100 fa '+(liked ? 'fa-heart' : 'fa-heart-o')"></i>
                                <i data-bs-toggle="modal" :data-bs-target="'#likesModal'+pid" class="my-comment-color ">
                                    {{likes.length}}
                                </i>
                            </button>
                            <!--comment-->
                            <button class="my-trnsp-btn m-0 p-1  d-flex justify-content-center align-items-center" 
							data-bs-toggle="modal" :data-bs-target="'#commentModal'+pid">
                                <i class="my-comment-color fa-regular fa-comment me-1"></i>
                                <i class="my-comment-color-2"> {{comments.length}}</i>
                            </button>
                        </div>
                        <div class="d-flex flex-row justify-content-start align-items-center ">
                            <p> Uploaded on <b>{{date}}</b></p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.my-card {
    width: 100%;
    max-width: 500px;
    border-radius: 10px;
    border: 1px solid #e0e0e0;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.photo-background-color {
    background-color: #f0f0f0;
    border-radius: 10px 10px 0 0;
}

.my-heart-color {
    color: #ff0000;
}


</style>