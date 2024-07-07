<script>
import CommentModal from '../components/CommentModal.vue'
import SuccessMsg from '../components/SuccessMsg.vue'

export default {

    components: {
        CommentModal,
        SuccessMsg
    },

    data: function() {
        return {
            errorMsg: null,
            successMsg: null,   
            Username: localStorage.getItem('username'),
            Uid: localStorage.getItem('token'),
            loading: false,
            image: null,
            images: null,
            comment: '',

            // Comments
            Comments : {
                comments : [
                    {
                        Uid : '',
                        Pid : '',
                        Message : '',
                    }
                ],
            },
            Message: '',

            // Stream
            Stream: {
                photos: [
                    {
                        Pid: '',
                        Uid: '',
                        File: '',
                        Date: '',
                    }
                ],
            },

            // Like
            Like: {
                Uid: '',
                Pid: '',
            },

            // Profile
            Profile: {
                Uid: '',
                Username: '',
                Followers: 0,
                Followees: 0,
                PhotosCount: 0,
            },

            // Get User by Username
            getUser: ''
        }
    },

    methods: {

        async refresh() {
            await this.getMyStream()
        },

        async uploadPhoto() {
            this.images = this.$refs.image.files[0]
        },

        async submitPhoto() {
            if (this.photo == null) {
                this.errorMsg = "Please select a photo"
                return
            } else {
                try {
                    let response = await this.$axios.post("/photos", {
                        File: this.photo,
                    }, {
                        headers: {
                            Authorization: "Bearer " + this.Uid
                        }
                    })
                    this.successMsg = response.data.message

                } catch (error) {
                    if (error.response) {
                        this.errorMsg = error.response.data.message
                    }
                }
            }
        },

        async getUserPhotos(){
            try {
                const response = await this.$axios.get(`/photos/`, {
                    params: { Uid: this.$route.params.Uid },
                    headers: {
                        'Authorization': `Bearer ${Uid}`,
                        'Accept': 'application/json',
                    },
                });
                this.Photos = response.data
            } catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        },

        async getMyStream() {
            try {
                let response = await this.$axios.get("/users/" + this.Uid + "/stream", {
                    headers: {
                        Authorization: "Bearer " + this.Uid
                    }
                })
                this.Stream = response.data
                for (let i = 0; i < this.Stream.photos.length; i++) {
                    this.Stream.photos[i].File = "data:image/jpeg;base64," + this.Stream.photos[i].File
                }
            } catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        },

        async getUserProfile() {
            try {
                let response = await this.$axios.get("/users/" + this.$route.params.Uid+ "/profile", {
                    headers: {
                        Authorization: "Bearer " + this.Uid
                    }
                })
                this.Profile = response.data
            } catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        }, 

        async commentPhoto(Uid, Pid, Message) {
            if (Message == ""){
                this.errorMsg = "Comment cannot be empty"
                return
            }
            try {
                let response = await this.$axios.post("/photos/" + Pid + "/comments/" + Uid, {
                    Message: Message,
                }, {
                    headers: {
                        Authorization: "Bearer " + this.Uid
                    }
                })
                this.clear = response.data
                this.refresh()
            } catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        },

        async openComments(Pid) {
            try {
                let response = await this.$axios.get("/photos/" + Pid + "/comments", {
                    headers: {
                        Authorization: "Bearer " + this.Uid
                    }
                })
                this.Comments = response.data
                const modal = this.$refs.commentModal
                modal.show()
            } catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        },

        async likePhoto(Pid, Uid) {
            try {
                let response = await this.$axios.put("/photos/" + Pid + "/likes/" + Uid, {
                    headers: {
                        Authorization: "Bearer " + this.Uid
                    }
                })
                this.clear = response.data
                this.refresh()
            } catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        },

        async unlikePhoto(Pid, Uid) {
            try {
                let response = await this.$axios.delete("/photos/" + Pid + "/likes/" + Uid, {
                    headers: {
                        Authorization: "Bearer " + this.Uid
                    }
                })
                this.clear = response.data
                this.refresh()
            } catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        },

        async getLikes(){
            try {
                let response = await this.$axios.get("/photos/" + this.$route.params.Pid + "/likes", {
                    headers: {
                        Authorization: "Bearer " + this.Uid
                    }
                })
                this.likes = response.data
            } catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        },

        searchUserUsername: "",
        async SearchUser() {
			if (this.searchUserUsername === this.username) {
				this.errormsg = "You can't search yourself."
			} else if (this.searchUserUsername === "") {
				this.errormsg = "Emtpy username field."
			} else {
				try {
					let response = await this.$axios.get("users/" + this.searchUserUsername + "/profile", {
						headers: {
							Authorization: "Bearer " + localStorage.getItem("token")
						}
					})
					this.profile = response.data
					this.$router.push({ path: '/users/' + this.searchUserUsername + '/view' })
				} catch (e) {
					if (e.response && e.response.status === 400) {
						this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
						this.detailedmsg = null;
					} else if (e.response && e.response.status === 500) {
						this.errormsg = "User does not exist on WASAPhoto.";
						this.detailedmsg = e.toString();
					} else {
						this.errormsg = e.toString();
						this.detailedmsg = null;
					}
				}
			}
		},

        async ViewProfile() {
            this.$router.push({path: '/users/' + this.Uid + '/profile'})
        },

        async Logout() {
            localStorage.removeItem('username')
            localStorage.removeItem('Uid')
            this.$router.push('/')
        },
    },

    mounted() {
        this.getMyStream()
    }
}
</script>

<template>
	<div>
		<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
			<div class="position-sticky pt-3 sidebar-sticky">
				<h6
					class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
					<span>General</span>
				</h6>
				<ul class="nav flex-column">
					<li class="nav-item">
						<RouterLink to="/session" class="nav-link">
							<svg class="feather">
								<use href="/feather-sprite-v4.29.0.svg#home" />
							</svg>
							Home
						</RouterLink>
					</li>
				</ul>
			</div>
		</nav>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Welcome back {{this.username }}</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button class="btn btn-danger" type="button" @click="Logout">Logout</button>
					<button class="btn btn-primary" type="button" @click="ViewProfile">Profile</button>
					<input type="file" accept="image/*" class="btn btn-outline-primary" @change="uploadPhoto" ref="file">
					<button class="btn btn-success" @click="submitPhoto">Upload</button>
				</div>
			</div>
		</div>
		<div class="input-group mb-3">

			<input type="text" id="searchUserUsername" v-model="searchUserUsername" class="form-control"
				placeholder="Search a user in WASAPhoto." aria-label="Recipient's username"
				aria-describedby="basic-addon2">
			<div class="input-group-append">
				<button class="btn btn-primary" type="button" @click="SearchUser">Search</button>
			</div>
		</div>

		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<SuccessMsg v-if="successmsg" :msg="successmsg"></SuccessMsg>

		<LogModal id="logviewer" :log="Comments" :token="token"></LogModal>


		<div class="row">
			<div class="col-md-4" v-for="photo in Stream.photos" :key="photo.Pid">
				<div class="card mb-4 shadow-sm">
					<img class="card-img-top" :src=photo.File alt="Card image cap">
					<div class="card-body">
						<RouterLink :to="'/users/' + photo.Uid + '/view'" class="nav-link">
							<button type="button" class="btn btn-outline-primary">{{photo.Uid}}</button>
						</RouterLink>
						<div
							class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
						</div>
						<p class="card-text">Uploaded on : {{photo.Date}}</p>

						<div class="input-group mb-3">
							<input type="text" id="comment" v-model="comment" class="form-control"
								placeholder="Comment!" aria-label="Recipient's username"
								aria-describedby="basic-addon2">
							<div class="input-group-append">
								<button class="btn btn-primary" type="button"
									@click="commentPhoto(photo.Uid, photo.Pid, photo.Message)">Send</button>
							</div>
						</div>

						<div class="d-flex justify-content-between align-items-center">
							<div class="btn-group">
								<button type="button" class="btn btn-dark"
									@click="openComments(photo.Pid)">Comments</button>
								<button type="button"  class="btn btn-primary"
									@click="likePhoto(photo.Pid, photo.Uid)">Like</button>
								<button type="button"  class="btn btn-danger"
									@click="unlikePhoto(photo.Pid, photo.Uid)">Unlike</button>
							</div>
						</div>


					</div>
				</div>
			</div>
		</div>
	</div>
</template>