<script>
const token = sessionStorage.getItem('token');
export default {

    props: ['photoId'],

    data() {
        return {
            showModal: false,
            comments: [],
            token: sessionStorage.getItem('token'),
        };
    },

    async created() {
        this.fetchComments();
    },

    methods: {

        async fetchComments() {
            try {
                const response = await this.$axios.get(`/photos/${this.photoId}/comments`, {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });
                console.log(token);
                console.log('Response:', response.data);
                this.comments = response.data;
                console.log('Comments:', this.comments);
            } catch (error) {
                if (error.response) {
                    const statusCode = error.response.status;
                    switch (statusCode) {
                        case 401:
                            console.error('Access Unauthorized:', error.response.data);
                            break;
                        case 403:
                            console.error('Access Forbidden:', error.response.data);
                            break;
                        case 404:
                            console.error('Not Found:', error.response.data);
                            break;
                        default:
                            console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
                    }
                } else {
                    console.error('Error:', error);
                }
            }

        },

        async deleteComment(pid, cid) {
            try {
                const response = await this.$axios.delete(`/photos/${pid}/comments/${cid}`, {
                    headers: {
                        'Authorization': `Bearer ${this.token}`,
                    }
                },);
                location.reload();
            }
            catch (error) {
                console.error(error, "cant delete!")
            }

        },
    },
};
</script>

<template>
    <div class="modal fade" tabindex="-1" :id="'listModal' + this.photoId" aria-labelledby="ModalLabel" aria-hidden="true">
        <div class="modal-dialog" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">Comments</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <ul class="list-group">
                        <li v-for="comment in comments" :key="comment.pid" class="list-group-item">
                            <div>
                                <strong>{{ comment.uid }} </strong>
                            </div>
                            <div>{{ comment.message }}</div>
                            <div v-if="comment.uid == this.token">
                                <button @click="deleteComment(comment.pid, comment.cid)" class="btn btn-danger btn-sm">Delete</button>
                            </div>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</template>