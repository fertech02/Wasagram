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