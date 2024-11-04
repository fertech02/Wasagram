<script>
const token = sessionStorage.getItem('authToken');

export default {
    data() {
        return {
            searchQuery: '',
            searchExecuted: false,
            Text: '',
            UserList: [],
        };
    },
    methods: {
        async searchUsers() {
            try {
                console.log("search started")
                const response = await this.$axios.get(`/users/`, {
                    params: { userName: this.searchQuery },
                    headers: {
                        'Authorization': `Bearer ${token}`,
                        'Accept': 'application/json',
                    },
                });
                console.log("search finished")
                this.searchExecuted = true;
                this.UserList = response.data.UList;
                this.Text = this.UserList === null ? "No user found with that name." : "";
            }
            catch (error) {
                console.error(error, "Error in user search")
                this.searchExecuted = true;
                if (error.response) {
                    const statusCode = error.response.status;
                    this.notBanned = false;
                    switch (statusCode) {
                        case 401:
                            console.error('Access Unauthorized:', error.response.data);
                            // unauthorized
                            this.Text = "You have to log in first";
                            break;
                        case 403:
                            console.error('Access Forbidden:', error.response.data);
                            // forbidden
                            break;
                        case 404:
                            console.error('Not Found:', error.response.data);
                            // not found
                            this.Text = "No users with such username";

                            break;
                        default:
                            console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
                            this.Text = "No users with such username";
                    }
                } else {
                    console.error('Error:', error);
                }
            }
        },
    },
};
</script>
