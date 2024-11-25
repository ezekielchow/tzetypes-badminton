<template>
  <div class="game-dashboard">
    <h1>Game Dashboard</h1>
    <p>list of all games.</p>

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
          <th @click="setSortKey('club_id')">Club ID</th>
          <th @click="setSortKey('game_type')">Game Type</th>
          <th @click="setSortKey('createdAt')">Created At</th>
          <th @click="setSortKey('isEnded')">Ended</th>
        </tr>
      </thead>
      <tbody>
       <tr
          v-for="game in sortedGames"
          :key="game.id"
          :class="{ ongoing: !game.isEnded, ended: game.isEnded }"
          @click="goToStatistics(game.id)"
        >
          <td>{{ game.club_id }}</td>
          <td>{{ game.game_type }}</td>
          <td>{{ game.serving_side }}</td>
          <td>
            {{ game.left_odd_player_name }}, 
            {{ game.left_even_player_name }}
          <span class="vs"> vs </span>
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
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router'; // For navigation

// fake data
const games = ref([
  {
    id: '1',
    club_id: 'clubowner-123',
    game_type: 'Singles',
    serving_side: 'Left',
    createdAt: '2024-11-01,10:00:00',
    isEnded: false,
    left_odd_player_name: 'Ahmad',
    left_even_player_name: 'Abu',
    right_odd_player_name: 'Jimmy',
    right_even_player_name: 'Long',
  },
  {
    id: '2',
    club_id: 'clubowner-124',
    game_type: 'Doubles',
    serving_side: 'Right',
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
const sortKey = ref('createdAt'); // default sort
const sortOrder = ref(1); // set 1 for ascending
const router = useRouter();

// format date
const formatDate = (dateString: string) => {
  const date = new Date(dateString);
  return date.toLocaleDateString();
};

// filter games based on search query for player names
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

// sort games based on sortKey
const sortedGames = computed(() => {
  return filteredGames.value.slice().sort((a, b) => {
    const aKey = a[sortKey.value];
    const bKey = b[sortKey.value];

    if (aKey < bKey) return -1 * sortOrder.value;
    if (aKey > bKey) return 1 * sortOrder.value;
    return 0;
  });
});

// set sort key and toggle sort order
const setSortKey = (key: string) => {
  if (sortKey.value === key) {
    sortOrder.value *= -1; // toggle order
  } else {
    sortKey.value = key;
    sortOrder.value = 1; // set ascending to default
  }
};

// navigate to statistics screen
const goToStatistics = (gameId: string) => {
  router.push(`/statistics/${gameId}`); 
};
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

.vs {
  font-weight: bold;
  font-size: 1.2rem;
  margin: 0 15px;
  align-items: center;
  text-transform: uppercase;
}
</style>
