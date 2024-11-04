<script>
export default {

    props: {
        photoId: String,
    },

    data() {
        return {
            showModal: false,
            comments: [],
            token: sessionStorage.getItem('authToken'),
        };
    },

    async created() {
        this.fetchComments();
    },

    methods: {

        async fetchComments() {
            try {
                const response = await this.$axios.get(`/photos/${this.photoId}/comments/`, {
                    headers: {
                        Authorization: `Bearer ${this.token}`,
                    },
                });
                console.log('Response:', response.data);
                this.comments = response.data.CList;
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

        async deleteComment(pId, cId) {
            try {
                const response = await this.$axios.delete(`/photos/${pId}/comments/${cId}`, {
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