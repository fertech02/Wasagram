<script>
import PhotoCard from '@/components/PhotoCard.vue';
const token = sessionStorage.getItem('token');

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
            isbanned: false,
            isfollowed: false,
            isItMe: false,
            photoList: [],
            reloadFlag: true,
        };
    },

    watch: {
        '$route.params.uid'(newParam, oldParam) {
            if (newParam !== oldParam) {
                this.refresh();
            }
        },
    },

    async created() {
        const uid= this.$route.params.uid;
        const token = sessionStorage.getItem('token');
        this.isItMe = (uid == token);
        this.fetchUserData();
    },

    methods: {

        refresh() {
            location.reload();
        },
        
        async fetchUserData() {
            const uid = this.$route.params.uid;
            try {
                const response = await this.$axios.get(`/users/${uid}/profile`, {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });
                this.found = true;
                this.username = response.data.username;
                this.followCount = response.data.followCount;
                this.followedCount = response.data.followedCount;
                this.photoCount = response.data.photoCount;
                this.isbanned = response.data.isBanned;
                this.isfollowed = response.data.isFollowed;
                this.photoList = response.data.photoList;

                console.log("PhotoList: ",this.photoList);

            } catch (error) {
                if (error.response) {
                    const statusCode = error.response.status;
                    switch (statusCode) {
                        case 400:
                            console.error('Bad request');
                            this.username = "You have to login first"
                        case 401:
                            console.error('Access Unauthorized:', error.response.data);
                            // unauthorized
                            this.userName = "You are not logged in"
                            break;
                        case 403:
                            console.error('Access Forbidden:', error.response.data);
                            // forbidden
                            this.username = "You have been banned by the user"
                            break;
                        case 404:
                            console.error('Not Found:', error.response.data);
                            // not found
                            if (uid === "null") {
                                this.username = "You are not logged in";
                            }
                            else {
                                this.username = "User not found";
                            }
                            break;
                        default:
                            console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
                    }
                } else {
                    console.error('Error:', error.message);
                }
            }
        },

        async Follow() {
            this.isfollowed = !this.isfollowed;
            const uid = this.$route.params.uid;
            try {
                if (this.isfollowed) {
                    this.followCount += 1;
                    await this.$axios.put(`/users/${token}/follow/${uid}`, {
                    }, {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    });
                } else {
                    this.followCount -= 1;
                    await this.$axios.delete(`/users/${token}/follow/${uid}`, {
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
            this.isbanned = !this.isbanned;
            const uid = this.$route.params.uid;
            try {
                if (this.isbanned) {
                    await this.$axios.put(`/users/${token}/ban/${uid}`, {
                    }, {
                        headers: {
                            Authorization: `Bearer ${token}`
                        }
                    });
                } else {
                    await this.$axios.delete(`/users/${token}/ban/${uid}`, {
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
                            {{ isfollowed ? 'Unfollow' : 'Follow' }} <svg class="feather">
                                <use href="/feather-sprite-v4.29.0.svg#user-plus" />
                            </svg>
                        </button>
                        <button @click="Ban" class="btn btn-danger">
                            {{ isbanned ? 'Unban' : 'Ban' }} <svg class="feather">
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
            <PhotoCard v-for="photo in photoList" :key="photo.pid" :photoId="photo.pid" :userId="photo.uid" :date="photo.date"  />
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

</style>