
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
            username: '',
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
        '$route.params.Uid'(newParam, oldParam) {
            if (newParam !== oldParam) {
                this.refresh();
            }
        },
    },
    async created() {
        const Uid = this.$route.params.Uid;
        this.isItMe = (Uid == token);
        this.getUserProfile();
    },
    methods: {
        refresh() {
            location.reload();
        },
        
        async getUserProfile() {
            try {
                    let response = await this.$axios.get("/users/" + Uid + "/profile", {
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
            const Uid = this.$route.params.Uid;
            const token = sessionStorage.getItem('authToken');
            try {
                if (this.isFollowed) {
                    this.followCount += 1;
                    await this.$axios.put(`/users/${token}/follow/${Uid}`, {
                    }, {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    });
                } else {
                    this.followCount -= 1;
                    await this.$axios.delete(`/users/${token}/follow/${Uid}`, {
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
            const Uid = this.$route.params.Uid;
            const token = sessionStorage.getItem('authToken');
            try {
                if (this.isBanned) {
                    await this.$axios.put(`/users/${token}/ban/${Uid}`, {
                    }, {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    });
                } else {
                    await this.$axios.delete(`/users/${token}/ban/${Uid}`, {
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

<template>
    <div class="container mt-5">
        <h1 class="display-4" style="font-size: 50px;">{{ username }}</h1>
        <div v-if="found">
            <div>
                <div v-if="!isItMe">
                    <div class="btn-group mt-1">
                        <button @click="Follow" class="btn btn-warning">
                            {{ isFollowed ? 'Unfollow' : 'Follow' }} <svg class="feather">
                                <use href="/feather-sprite-v4.29.0.svg#user-plus" />
                            </svg>
                        </button>
                        <button @click="Ban" class="btn btn-danger">
                            {{ isBanned ? 'Unban' : 'Ban' }} <svg class="feather">
                                <use href="/feather-sprite-v4.29.0.svg#slash" />
                            </svg>
                        </button>
                    </div>
                </div>
            </div>
            <div style="font-size: 20px;" class="container mt-2">
                <div class="row bg-light p-4 shadow-lg">
                    <div class="row">Followers: {{ followCount }}</div>
                    <div class="row">Followed: {{ followedCount }}</div>
                    <div class="row">Photos: {{ photoCount }}</div>
                </div>
            </div>

        </div>
        <hr />
        <div class="photos">
            <PhotoCard v-for="photo in photoList" :key="photo.Pid" :photoId="photo.Pid" :date="photo.date"
                :authorName="photo.Uid" :likeCount="photo.likecount" :caption="photo.caption" />
        </div>
    </div>
</template>
  
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