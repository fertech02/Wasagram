<script>
import CommentModal from '@/components/CommentModal.vue';
import CommentListModal from '@/components/CommentListModal.vue';

const token = sessionStorage.getItem('authToken');
export default {
  components: {
    CommentModal,
    CommentListModal,
  },
  props: {
    photoId: String,
    likeCount: Number,
    authorName: String,
    caption: String,
    date: String,
  },
  data() {
    return {
      imgSrc: null,
      isLiked: false,
      LikeCount: this.likeCount,
      authorId: "",
      isMe: false,
      notBanned: true,
      modalId: String(this.photoId),
    };
  },

  async created() {
    if (this.photoId) {
      try {
        const response = await this.$axios.get(`/photos/${this.photoId}`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
          responseType: 'blob',
        });
        const imageUrl = URL.createObjectURL(response.data);
        this.imgSrc = imageUrl;
        const isL = await this.$axios.get(`/photos/${this.photoId}/likes/${token}`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.isLiked = isL.data.isLiked
        this.findAuthorId();
      } catch (error) {
        if (error.response) {
          const statusCode = error.response.status;
          this.notBanned = false;
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
    }
  },

  computed: {

  },


  methods: {

    async findAuthorId() {
      try {
        const userId = this.$route.params.userId;
        const hasStreamSegment = this.$route.path.includes('/stream');
        if (userId == token && !hasStreamSegment) {
          this.isMe = true;
        };
      }
      catch (error) {
        console.error(error, "Error searching photo owner.")
      }
    },

    async deletePhoto() {
      try {
        const response = await this.$axios.delete(`/photos/${this.photoId}`, {
          headers: {
            'Authorization': `Bearer ${token}`,
          }
        },);
        location.reload();
      }
      catch (error) {
        console.error(error, "cant delete photo!")
      }
    },

    async likePhoto() {

      this.isLiked = !this.isLiked;
      
      try {
        const token = sessionStorage.getItem('authToken');
        if (this.isLiked) {
          this.LikeCount += 1;
          await this.$axios.put(`/photos/${this.photoId}/likes/${token}`, {
          }, {
            headers: {
              Authorization: `Bearer ${token}`
            }
          });
        } else {
          this.LikeCount -= 1;
          await this.$axios.delete(`/photos/${this.photoId}/likes/${token}`, {
            headers: {
              Authorization: `Bearer ${token}`
            }
          });

        }
      } catch (error) {
        console.error(error, "Error modifying like status.")
      }

    },
  },
  
};
</script>