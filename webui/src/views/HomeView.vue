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
            token: localStorage.getItem('token'),
            Username: localStorage.getItem('username'),
            Uid: localStorage.getItem('Uid'),
            loading: false,
            image: null,
            images: null,

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
                            Authorization: "Bearer " + this.token
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
                        'Authorization': `Bearer ${token}`,
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
                        Authorization: "Bearer " + this.token
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
                        Authorization: "Bearer " + this.token
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
                        Authorization: "Bearer " + this.token
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

        async openComments() {
            try {
                let response = await this.$axios.get("/photos/" + Pid + "/comments", {
                    headers: {
                        Authorization: "Bearer " + this.token
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
                        Authorization: "Bearer " + this.token
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
                        Authorization: "Bearer " + this.token
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
                        Authorization: "Bearer " + this.token
                    }
                })
                this.likes = response.data
            } catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        },

        async Logout() {
            localStorage.removeItem('token')
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
        <div class="container">
            <div class="row">
                <div class="col-md-12">
                    <div class="card">
                        <div class="card-header">
                            <h4>Home</h4>
                        </div>
                        <div class="card-body">
                            <div class="row">
                                <div class="col-md-12">
                                    <div class="form-group">
                                        <label for="image">Upload Photo</label>
                                        <input type="file" ref="image" class="form-control" @change="uploadPhoto">
                                    </div>
                                    <button class="btn btn-primary" @click="submitPhoto">Submit</button>
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-md-12">
                                    <div class="alert alert-danger" v-if="errorMsg" role="alert">
                                        {{ errorMsg }}
                                    </div>
                                    <div class="alert alert-success" v-if="successMsg" role="alert">
                                        {{ successMsg }}
                                    </div>
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-md-12">
                                    <div class="card">
                                        <div class="card-header">
                                            <h4>Stream</h4>
                                        </div>
                                        <div class="card-body">
                                            <div class="row">
                                                <div class="col-md-12">
                                                    <div class="alert alert-danger" v-if="errorMsg" role="alert">
                                                        {{ errorMsg }}
                                                    </div>
                                                    <div class="alert alert-success" v-if="successMsg" role="alert">
                                                        {{ successMsg }}
                                                    </div>
                                                </div>
                                            </div>
                                            <div class="row">
                                                <div class="col-md-12">
                                                    <div class="card-columns">
                                                        <div class="card" v-for="photo in Stream.photos" :key="photo.Pid">
                                                            <img :src="photo.File" class="card-img-top" alt="...">
                                                            <div class="card-body">
                                                                <h5 class="card-title">Photo</h5>
                                                                <p class="card-text">Date: {{ photo.Date }}</p>
                                                                <button class="btn btn-primary" @click="openComments(photo.Pid)">Comments</button>
                                                                <button class="btn btn-primary" @click="likePhoto(photo.Pid, Uid)">Like</button>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <success-msg ref="successMsg" :successMsg="successMsg"></success-msg>
        <comment-modal ref="commentModal" :comments="Comments.comments" @comment="commentPhoto"></comment-modal>
    </div>
</template>