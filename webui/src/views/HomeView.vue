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
            username: localStorage.getItem('username'),
            loading: false,

            // Comments
            PhotoComments : {
                Uid : '',
                Pid : '',
                comments : [
                    {
                        Uid : '',
                        Pid : '',
                        comment : '',
                    }
                ],
            },
            comment: '',

            // Stream
            Stream: {

            },

            // Like
            Like: {

            },

            // Profile
            Profile: {

            },

            // Get User by Username
            getUser: ''
        }
    },
    methods: {

        async uploadPhoto() {
        },
        async submitPhoto() {
        },
        async searchUser() {
        },
        async openLog(Uid, Pid) {

        },
        async refresh() {
            this.getMyStream()
        },
        async getMyStream() {
        },
        async getUser() {
        },
        async getUserProfile() {
        },
        async commentPhoto() {
        },
        async likePhoto() {
        },
        async unlikePhoto() {
        },
        async Logout() {
        },
    },
    mounted() {
        this.getMyStream()
    }
}
</script>

<template>
    <div>
        <b-container>
            <b-row>
                <b-col>
                    <b-card-group deck>
                        <b-card v-for="log in logs" :key="log.photoId" :title="log.photoId" img-src="https://picsum.photos/600/300/?image=25" img-alt="Image" img-top>
                            <b-card-text>
                                {{ log.comment }}
                            </b-card-text>
                            <b-button @click="openLog(log)" variant="primary">View</b-button>
                        </b-card>
                    </b-card-group>
                </b-col>
            </b-row>
        </b-container>
        <CommentModal :log="selectedLog" :token="token" @uncomment="refresh" />
        <SuccessMsg :message="successMsg" />
    </div>
</template>