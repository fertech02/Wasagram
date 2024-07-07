<template>
    <div class="container mt-5">
        <h1 class="display-4" style="font-size: 50px;">{{ userName }}</h1>
        <div v-if="found">
            <div>
                <div v-if="!isItMe">
                    <div class="btn-group mt-1">
                        <button @click="toggleFollow" class="btn btn-warning">
                            {{ isFollowed ? 'Unfollow' : 'Follow' }} <svg class="feather">
                                <use href="/feather-sprite-v4.29.0.svg#user-plus" />
                            </svg>
                        </button>
                        <button @click="toggleBan" class="btn btn-danger">
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
            <PhotoCard v-for="photo in photoList" :key="photo.id" :photoId="photo.id" :date="photo.date"
                :authorName="userName" :likeCount="photo.likecount" :caption="photo.caption" />
        </div>
    </div>
</template>
  
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
        this.fetchUserData();
    },
    methods: {
        refresh() {
            location.reload();
        },
       async getUserProfile() {},
       
        async Follow() {
            // frontend
            this.isFollowed = !this.isFollowed;
            // backend
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
            // frontend
            this.isBanned = !this.isBanned;
            // backend
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