<template>
  <div class="game-dashboard">
    <h1>Game Dashboard</h1>
    <p>Here is the list of all games.</p>

    // Search box
    <div class="search-container">
      <input
        type="text"
        v-model="searchQuery"
        placeholder="Search by player name..."
        class="search-input"
      />
    </div>

    <table class="game-listing-table">
      <thead>
        <tr>
          <th>Game Name</th>
          <th>Status</th>
          <th>Players</th>
          <th>Created At</th>
          <th>Ended</th>
        </tr>
      </thead>
      <tbody>
        
        // looping through seeded data
        <tr
          v-for="game in filteredGames"
          :key="game.id"
          :class="{
            'ongoing': !game.isEnded,
            'ended': game.isEnded
          }"
        >
          <td>{{ game.name }}</td>
          <td>{{ game.status }}</td>
          <td>
            {{ game.left_odd_player_name }}, 
            {{ game.left_even_player_name }}, 
            {{ game.right_odd_player_name }}, 
            {{ game.right_even_player_name }}
          </td>
          <td>{{ formatDate(game.createdAt) }}</td>
          <td>{{ game.isEnded ? 'Yes' : 'No' }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';

// Seeded data
const games = ref([
  {
    id: 1,
    name: 'Game 1',
    status: 'Ongoing',
    createdAt: '2024-11-01,10:00:00',
    isEnded: false,
    left_odd_player_name: 'Ahmad',
    left_even_player_name: 'Abu',
    right_odd_player_name: 'Jimmy',
    right_even_player_name: 'Long',
  },
  {
    id: 2,
    name: 'Game 2',
    status: 'Completed',
    createdAt: '2024-11-02,11:30:00',
    isEnded: true,
    left_odd_player_name: 'Sing',
    left_even_player_name: 'Mei',
    right_odd_player_name: 'Pang',
    right_even_player_name: 'Hank',
  },
  
]);

// Search query for filtering games by player name
const searchQuery = ref('');

// Format date
const formatDate = (dateString: string) => {
  const date = new Date(dateString);
  return date.toLocaleDateString();
};

// Computed property to filter games based on the search query for player names
const filteredGames = computed(() => {
  return games.value.filter((game) => {
    const allPlayers = [
      game.left_odd_player_name,
      game.left_even_player_name,
      game.right_odd_player_name,
      game.right_even_player_name,
    ]
      .join(' ')
      .toLowerCase();

    return allPlayers.includes(searchQuery.value.toLowerCase());
  });
});

onMounted(() => {
  
});
</script>

<style scoped>
.game-dashboard {
  padding: 1rem;
}

.search-container {
  margin-bottom: 1rem;
}

.search-input {
  width: 100%;
  padding: 0.5rem;
  font-size: 1rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.game-listing-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 1rem;
}

.game-listing-table th,
.game-listing-table td {
  padding: 0.8rem;
  text-align: left;
  border: 1px solid #ddd;
}

.ongoing {
  background-color: #28a745;
  color: white;
}

.ended {
  background-color: #dc3545;
  color: white;
}
</style>
