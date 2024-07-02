<script>
import CommentModal from '../components/CommentModal.vue';

export default {

    mounted() {
        if (localStorage.getItem("reloadedStream")) {
            localStorage.removeItem("reloadedStream");
        } else {
            localStorage.setItem("reloadedStream", "1");
            location.reload();
        }
    },

    components: {
        CommentModal,
    },

    data: function() {
        return {
            
            // Profile
            Profile: {
                Uid: '',
                Username: '',
                Followers: 0,
                Followees: 0,
                PhotosCount: 0,
            },

            // Photos
            Photos: {
                photos: [
                    {
                        Pid: '',
                        Uid: '',
                        File: '',
                        Date: '',
                    }
                ],
            },

            // Follow
            Follow: {
                FolloweeId: '',
                FollowerId: '',
            },

            // Ban
            Ban: {
                BannerId: '',
                BannedId: '',
            },

            // Comments
            Comments: {
                comments: [
                    {
                        Uid: '',
                        Pid: '',
                        Message: '',
                    }
                ],
            },
            Message: '',

            // Get User by Username
            getUser: ''
        }
    },

    watch: {
        '$route.params.Uid' (newParam, oldParam) {
            if (newParam !== oldParam) {
                this.refresh();
            }
        }
    },

    mounted() {
        this.getUserProfile()
    },

    methods: {
        async refresh(){
            await this.getUserProfile()
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
        
        async banUser(bannerId, bannedId){
            try {
                let response = await this.$axios.post("/users/" + bannerId + "/ban/" + bannedId, {
                    headers: {
                        Authorization: "Bearer " + this.token
                    }
                })
                this.clear = response.data
                this.refresh()
            }
            catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        },

        async unbanUser(bannerId, bannedId){
            try{
                let response = await this.$axios.delete("/users/" + bannerId + "/ban/" + bannedId, {
                    headers: {
                        Authorization: "Bearer " + this.token
                    }
                })
                this.ban = response.data
            }
            catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        },

        async followUser(followeeId, followerId){
            try {
                let response = await this.$axios.post("/users/" + followeeId + "/follow/" + followerId, {
                    headers: {
                        Authorization: "Bearer " + this.token
                    }
                })
                this.clear = response.data
                this.refresh()
                this.successmsg = "User followed successfully"
            }
            catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        },

        async unfollowUser(followeeId, followerId){
            try {
                let response = await this.$axios.delete("/users/" + followeeId + "/follow/" + followerId, {
                    headers: {
                        Authorization: "Bearer " + this.token
                    }
                })
                this.clear = response.data
                this.refresh()
                this.successmsg = "User unfollowed successfully"
            } catch (error) {
                if (error.response) {
                    this.errorMsg = error.response.data.message
                }
            }
        },

        async commentPhoto(Uid, Pid, Message){
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
    }

};
</script>

<template>
    <div>
        <div class="container">
            <div class="row">
                <div class="col-md-4">
                    <div class="card">
                        <div class="card-body">
                            <div class="text-center">
                                <img :src="'/users/' + Profile.Uid + '/avatar'" class="rounded-circle" width="150" height="150">
                                <h3>{{ Profile.Username }}</h3>
                                <p>Followers: {{ Profile.Followers }}</p>
                                <p>Followees: {{ Profile.Followees }}</p>
                                <p>Photos: {{ Profile.PhotosCount }}</p>
                                <button v-if="Profile.Uid != user.Uid" v-on:click="followUser(Profile.Uid, user.Uid)" class="btn btn-primary">Follow</button>
                                <button v-else v-on:click="unfollowUser(Profile.Uid, user.Uid)" class="btn btn-primary">Unfollow</button>
                                <button v-if="Profile.Uid != user.Uid" v-on:click="banUser(user.Uid, Profile.Uid)" class="btn btn-danger">Ban</button>
                                <button v-else v-on:click="unbanUser(user.Uid, Profile.Uid)" class="btn btn-danger">Unban</button>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-md-8">
                    <div class="row">
                        <div class="col-md-4" v-for="photo in Photos.photos">
                            <div class="card">
                                <img :src="'/photos/' + photo.Pid + '/file'" class="card-img-top" alt="...">
                                <div class="card-body">
                                    <p class="card-text">{{ photo.Date }}</p>
                                    <button v-on:click="openComments(photo.Pid)" class="btn btn-primary">Comments</button>
                                    <button v-on:click="likePhoto(photo.Pid, user.Uid)" class="btn btn-primary">Like</button>
                                    <button v-on:click="unlikePhoto(photo.Pid, user.Uid)" class="btn btn-primary">Unlike</button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <comment-modal ref="commentModal" :comments="Comments.comments"></comment-modal>
    </div>
</template>