<script>
import CommentModal from '../components/CommentModal.vue'

export default {
    components: {
        CommentModal
    },
    data: function() {
        return {

            errorMsg: null,
            Uid: localStorage.getItem('Uid'),
            Username: localStorage.getItem('Username'),
            token: localStorage.getItem('token'),
            newUsername: '',

            // User Profile
            Profile : {
                Uid: '',
                Username: '',
                Followers: 0,
                Followees: 0,
                PhotosCount: 0,
            },

            // User Photos
            Photos : {
                photos: [
                    {
                        Pid: '',
                        Uid: '',
                        File: '',
                        Date: '',
                    }
                ],
            },

            // Photo Comments
            Message: "",
            Comments: {
                comments: [
                    {
                        Uid: '',
                        Pid: '',
                        Message: '',
                    }
                ],
            },
            

            // Followees
            Followees : [],
            
            // Followers
            Followers : [],

            // Get User by Username
            getUser: ''
        }
    },
    methods: {
        async refresh() {
            await this.getUserProfile(),
            await this.getUserPhotos()
        },

        async getUserProfile(){
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

        async deletePhoto(Pid){
            try {
                let response = await this.$axios.delete("/photos/" + Pid, {
                    headers: {
                        Authorization: "Bearer " + this.token
                    }
                })
                this.refresh()
            } catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        },

        async setMyUserName(){
            if (this.newUsername == ""){
                this.errorMsg = "Username cannot be empty"
                return
            }
            try {
                let response = await this.$axios.put("/users/" + this.Uid + "/username", {
                    Username: this.newUsername,
                }, {
                    headers: {
                        Authorization: "Bearer " + this.token
                    }
                })
            } catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        },

        async commentPhoto(Pid, Uid, Message){
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

        async openComments(){
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

        async likePhoto(Pid, Uid){
            try {
                let response = await this.$axios.post("/photos/" + Pid + "/likes/" + Uid, {
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

        async unlikePhoto(Pid, Uid){
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

        async getFollowedUsers(){
            try {
                let response = await this.$axios.get("/users/" + this.Uid + "/follow", {
                    headers: {
                        Authorization: "Bearer " + this.token
                    }
                })
                this.Followees = response.data
            } catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        },

        // async getFollowers(){}
    }, 
}
</script>

<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="col-md-4">
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title">Profile</h5>
                            <p class="card-text">Username: {{ Profile.Username }}</p>
                            <p class="card-text">Followers: {{ Profile.Followers }}</p>
                            <p class="card-text">Followees: {{ Profile.Followees }}</p>
                            <p class="card-text">Photos: {{ Profile.PhotosCount }}</p>
                            <button class="btn btn-primary" @click="setMyUserName">Change Username</button>
                            <input type="text" v-model="newUsername" placeholder="New Username">
                        </div>
                    </div>
                </div>
                <div class="col-md-8">
                    <div class="card">
                        <div class="card-body">
                            <h5 class="card-title
                            ">Photos</h5>
                            <div class="row">
                                <div class="col-md-4" v-for="photo in Photos.photos" :key="photo.Pid">
                                    <img :src="photo.File" class="img-fluid" alt="Responsive image">
                                    <button class="btn btn-primary" @click="deletePhoto(photo.Pid)">Delete</button>
                                    <button class="btn btn-primary" @click="openComments(photo.Pid)">Comments</button>
                                    <button class="btn btn-primary" @click="likePhoto(photo.Pid, Uid)">Like</button>
                                    <button class="btn btn-primary" @click="unlikePhoto(photo.Pid, Uid)">Unlike</button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <CommentModal ref="commentModal" :comments="Comments.comments" @comment="commentPhoto"></CommentModal>
    </div>
</template>
