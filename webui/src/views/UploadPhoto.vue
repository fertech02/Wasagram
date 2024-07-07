</template>


<script>
const token = sessionStorage.getItem('authToken');
export default {
  data() {
    return {
      photo: null,
      caption: '',
      uploadSuccess: false,
      endText: '',
    };
  },
  methods: {
    onFileChange(event) {
      this.photo = event.target.files[0];
    },
    async uploadPhoto() {
      if (!this.photo) {
        console.log('Photo is required');
        return;
      }

      const formData = new FormData();
      formData.append('file', this.photo);
      const additionalData = {
        Caption: this.caption,
      };

      formData.append('additionalData', JSON.stringify(additionalData));
      const config = {
        headers: {
          'Content-Type': 'multipart/form-data',
          'Authorization': `Bearer ${token}`,
        },
      };
      try {
        const response = await this.$axios.post(`/photos/`, formData, config);
        console.log('Photo uploaded successfully', response.data);
        this.endText = "Photo uploaded!";
        this.uploadSuccess = true;
      }
      catch (error) {
        const statusCode = error.response.status;
        switch (statusCode) {
          case 401:
            console.error('Access Unauthorized');
            this.endText = "You have to log in to post a photo";
            this.uploadSuccess = true;
            break;
          default:
            console.error(`Unhandled HTTP Error (${statusCode}):`, error.response.data);
        }
      }
    },
  },
};
</script>