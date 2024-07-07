
<script>
import PhotoCard from '@/components/PhotoCard.vue';
const token = sessionStorage.getItem('authToken');

export default {
    mounted() {
        if (localStorage.getItem('reloadedstream')) {
            localStorage.removeItem('reloadedstream');
        } else {
            localStorage.setItem('reloadedstream', '1');
            location.reload();
        }
    },
    data() {
        return {
            userName: '',
            found: false,
            followCount: 0,
            followedCount: 0,
            photoCount: 0,
            isBanned: false,
            isFollowed: false,
            isItMe: false,
            photoList: [],
            reloadFlag: true,

            profile: {
                user_id: '',
                username: '',
                followers: 0,
                followees: 0,
                photos: 0,
            },
        };
    },
    watch: {
        '$route.params.userId'(newParam, oldParam) {
            if (newParam !== oldParam) {
                this.refresh();
            }
        },
    },
    async created() {
        const userId = this.$route.params.userId;
        this.isItMe = (userId == token);
        this.getUserProfile();
    },
    methods: {
        refresh() {
            location.reload();
        },
        
        async getUserProfile() {
            try {
                    let response = await this.$axios.get("/users/" + this.username + "/profile", {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("token")
                        }
                    })
                    this.profile = response.data
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
        },
       
        async Follow() {
            this.isFollowed = !this.isFollowed;
            const userId = this.$route.params.userId;
            const token = sessionStorage.getItem('authToken');
            try {
                if (this.isFollowed) {
                    this.followCount += 1;
                    await this.$axios.put(`/users/${token}/follows/${userId}`, {
                    }, {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    });
                } else {
                    this.followCount -= 1;
                    await this.$axios.delete(`/users/${token}/follows/${userId}`, {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    });
                }
            } catch (error) {
                console.error(error, "Error modifying follow status.")
            }

        },
        async Ban() {
            this.isBanned = !this.isBanned
            const userId = this.$route.params.userId;
            const token = sessionStorage.getItem('authToken');
            try {
                if (this.isBanned) {
                    await this.$axios.put(`/users/${token}/bans/${userId}`, {
                    }, {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    });
                } else {
                    await this.$axios.delete(`/users/${token}/bans/${userId}`, {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    });

                }
            } catch (error) {
                console.error(error, "Error modifying ban status.")
            }
        },
    },
    components: {
        PhotoCard,
    },
};
</script>
  
<style scoped>
.user-info {
    text-align: center;
    font-size: 20px;
}

hr {
    margin: 20px 0;
}

.photos {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-around;
}

.photos .photo-card {
    margin-bottom: 30px;
}
</style>