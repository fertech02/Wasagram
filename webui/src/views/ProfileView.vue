<script>
// get user profile
export default {
	data: function() {
		return {
            errormsg: null,
            token: sessionStorage.getItem('token'),
			username: "",
            photosCount: 0,
            followersCount: 0,
            followingCount: 0,
            isOwner: false,
            doIFollowUser: false,
            isInMyBannedList: false,
            amIBanned: false,

            // getPhotosList
            photosList: [],

            // getFollowersList
            followerList: [],

            // getFollowingsList
            followingList: [],

            userExists: false,
            user_id: "",
		}
	},
    watch: {
        // property to watch
        pathUsername(newUName, oldUName) {
            if (newUName !== oldUName){
                this.getUserProfile()
            }
        }
    },
    computed: {
        pathUsername() {
            return this.$route.params.username
        },
    },
    methods: {

        async getUserProfile() {
            try {
                let response = await this.$axios.get("/users/" + this.token + "/profile/");
                let profile = response.data;
                this.photosCount = profile.photo_count;
                this.followersCount = profile.followers;
                this.followingCount = profile.followees;
                if (profile.user_id == this.token) {
                    this.isOwner = true;
                } else {
                    this.isOwner = false;
                }
            } catch (error) {
                const status = error.response.status;
                const reason = error.response.data;
                this.errormsg = `Status ${status}: ${reason}`;
            }
        },

		async Follow() {
            try {
                if (this.doIFollowUser) {
                    await this.$axios.delete("/users/" + this.token + "/follow/" + this.user_id );
                    this.doIFollowUser = false;
                    this.followersCount -= 1;
                    this.getUserProfile();
                } else {
                    await this.$axios.post("/users/" + this.token + "/follow/" + this.user_id);
                    this.doIFollowUser = true;
                    this.followersCount += 1;
                    this.getUserProfile();
                }
            }
            catch (error) {
                const status = error.response.status;
                const reason = error.response.data;
            }
        },

        async Ban() {
            try {
                if (this.isInMyBannedList) {
                    await this.$axios.delete("/users/" + this.token + "/ban/" + this.user_id);
                    this.isInMyBannedList = false;
                    this.getUserProfile();
                } else {
                    await this.$axios.post("/users/" + this.token + "/ban/" + this.user_id);
                    this.isInMyBannedList = true;
                    this.getUserProfile();
                }
            }
            catch (error) {
                const status = error.response.status;
                const reason = error.response.data;
            }
		},

        async uploadPhoto() {
            try {
                let file = document.getElementById("fileUploader").files[0];
                const reader = new FileReader();
                reader.onload = async (e) => {
                    let response = await this.$axios.post("/photos/", reader.result);
                    this.photosList.unshift(response.data);
                    this.photosCount += 1;
                }
            }  
            catch (error) {
                const status = error.response.status;
                const reason = error.response.data;
            }
        },

        async getPhotosList() {
            try {
                let response = await this.$axios.get("/photos", {params: {user_id: this.user_id}});
                this.photosList = response.data === null ? [] : response.data;
            } catch (error) {
				const status = error.response.status;
        		const reason = error.response.data;
            }
        },

        async getFollowingList() {
            try {
                let response = await this.$axios.get("/users/" + this.token + "/follow/");
                this.followerList = response.data;
                if (this.followerList == null) {
                    this.followerList = [];
                }
            }
            catch (error) {
                const status = error.response.status;
                const reason = error.response.data;
            }
        },

        removePhotoFromList(pid){
            try {
                this.photosList = this.photosList.filter(photo => photo.Pid !== pid);
                this.photosCount -= 1;
            }
            catch (error) {
                const status = error.response.status;
                const reason = error.response.data;
            }
		},

        visitUser(token) {
            this.$router.push(`/users/${token}/profile`);
        }
    },
    mounted() {
        this.getUserProfile();
    }
}
</script>

<template>

    <UserModal
    :modalID="'usersModalFollowers'" 
    :usersList="followerList"
    @visitUser="visitUser"
    />

    <UserModal
    :modalID="'usersModalFollowing'" 
    :usersList="followingList"
    @visitUser="visitUser"
    />

    <div class="container-fluid" v-if="userExists && !amIBanned">
        <div class="row">
            <div class="col-12 d-flex justify-content-center">
                <div class="card w-50 container-fluid">
                    <div class="row">
                        <div class="col">
                            <div class="card-body d-flex justify-content-between align-items-center">
                                <h5 class="card-title p-0 me-auto mt-auto">@{{username}}</h5>

                                <button v-if="!isOwner && !isInMyBannedList" @click="Follow" class="btn btn-success ms-2">
                                    {{doIFollowUser ? "Unfollow" : "Follow"}}
                                </button>

                                <button v-if="!isOwner" @click="Ban" class="btn btn-danger ms-2">
                                    {{isInMyBannedList ? "Unban" : "Ban"}}
                                </button>
                            </div>
                        </div>
                    </div>

                    <div v-if="!isInMyBannedList" class="row mt-1 mb-1">
                        <button class="col-4 d-flex justify-content-center btn-foll">
                            <h6 class="ms-3 p-0 ">Posts: {{photosCount}}</h6>
                        </button>
                    
                        <button class="col-4 d-flex justify-content-center btn-foll">
                            <h6 data-bs-toggle="modal" :data-bs-target="'#usersModalFollowers'">
                                Followers: {{followersCount}}
                            </h6>
                        </button>
                    
                        <button class="col-4 d-flex justify-content-center btn-foll">
                            <h6 data-bs-toggle="modal" :data-bs-target="'#usersModalFollowing'">
                                Following: {{followingCount}}
                            </h6>
                        </button>
                    </div>
                </div>
            </div>
        </div>


        <div class="row">
            <div class="container-fluid mt-3">
                <div class="row ">
                    <div class="col-12 d-flex justify-content-center">
                        <h2>Posts</h2>
                        <input id="fileUploader" type="file" class="profile-file-upload" @change="uploadPhoto" accept=".jpg, .png">
                        <label v-if="isOwner" class="btn my-btn-add-photo ms-2 d-flex align-items-center" for="fileUploader"> Add </label>
                    </div>
                </div>
                <div class="row ">
                    <div class="col-3"></div>
                    <div class="col-6">
                        <hr class="border border-dark">
                    </div>
                    <div class="col-3"></div>
                </div>
            </div>
        </div>

        <div class="row">
            <div class="col">
                <div v-if="!isInMyBannedList && photosCount>0">
                    <Photo v-for="photo in photosList"
                    :key="photo.pid"
                    :pid="photo.pid"
                    :ownerID="photo.user_id"
                    :username="photo.username"
                    :date="photo.date"
                    :likesListParent="photo.likes"
                    :commentsListParent="photo.comments"
                    :isOwner="isOwner"
                    @removePhoto="removePhotoFromList"
                    />
                </div>
                
                <div v-if="!isInMyBannedList && photosCount==0" class="mt-5 ">
                    <h2 class="d-flex justify-content-center" style="color: white;">No posts yet</h2>
                </div>
            </div>
        </div>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    
</template>