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
      Pid: String,
      Uid: String,
      File: String,
      date: String,
    },
    data() {
      return {
        imgSrc: this.File,
        authorId: this.Uid,
        isLiked: false,
        isMe: false,
        notBanned: true,
        caption: '',
        LikeCount: 0,
        isMe: false,
      };
    },
  
    async created() {
      if (this.Pid) {
        try {
          const response = await this.$axios.get('/photos/' + this.Pid, {
            headers: {
              Authorization: `Bearer ${token}`,
            },
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
        }
        catch (error) {
          console.error(error);
        }
      } else {
        console.error('Error:', error)
      }
    },
    computed: {
  
    },
    methods: {

      async findAuthor() {
        try {
          const response = this.$route.params.uid;
          const hasStream = this.$route.path.includes('/stream');
          if (userId == token && !hasStream) {
            this.isMe = true;
          }

        } catch (error) {
          console.error(error);
        }
      },

      async deletePhoto(){
        try {
          const response = await this.$axios.delete(`/photos/${this.Pid}`, {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          });
          location.reload();
        } catch (error) {
          console.error(error);
        }
      },

      async likePhoto() {
        this.isLiked = !this.isLiked;
        try {
          if (this.isLiked) {
            this.LikeCount += 1;
            const response = await this.$axios.put('/photos/' + this.Pid + '/likes' + this.Uid, {}, {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          });
          } else {
            this.LikeCount -= 1;
            await this.$axios.delete('/photos/' + this.Pid + '/likes' + this.Uid, {
              headers: {
                Authorization: `Bearer ${token}`,
              },
            });
          }

        } catch (error) {
          console.error(error);
        }
      },
    }
  };
  </script>

<template>
  <div class="center-container">
    <div class="photo-card  ">
      <img :src="imgSrc" alt="Photo" width="100%" />
      <div class="photo-details">
        <div class="author">
          <router-link :to="'/users/' + authorId">{{ authorId }}</router-link>
        </div>
        <div class="actions">
          <div>
            <button v-on:click="likePhoto">Like</button>
            <div class="like-counter">{{ LikeCount }}</div>
          </div>
          <div>
            <button v-on:click="deletePhoto" v-if="isMe">Delete</button>
          </div>
        </div>
        <div class="caption">
          <div class="caption-border"></div>
          <div class="caption-text">{{ caption }}</div>
          <div class="caption-border"></div>
        </div>
      </div>
    </div>
  </div>
</template>
  
<style scoped>
  .center-container {
    display: flex;
    justify-content: center;
    align-items: center;
  }
  
  .photo-card {
    border: 3px solid #6d6969;
    border-radius: 4px;
    padding: 10px;
    width: 500px;
    text-align: center;
    font-family: 'Arial', sans-serif;
  }
  
  .photo-details {
    margin-top: 10px;
  }
  
  .author {
    font-size: 20px;
    margin-bottom: 5px;
  }
  
  .actions {
    display: flex;
    justify-content: space-between;
    margin: 15px;
  }
  
  .like-counter {
    margin-left: 2px;
    border: 2px solid #d102027a;
    border-radius: 4px;
    padding: 8px;
  }
  
  .caption {
    display: flex;
    align-items: center;
    margin-top: 10px;
  }
  
  .caption-border {
    flex: 1;
    height: 3px;
    background-color: #1a1212;
    padding: 4px;
    margin-top: 10px;
    margin-bottom: 10px;
  
  }
  
  .caption-text {
    padding: 0 10px;
  }
  </style>