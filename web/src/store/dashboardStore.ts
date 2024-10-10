import { MyApi } from '@/services/requests';
import { defineStore } from 'pinia';

export const useDashboardStore = defineStore('dashboard', {
  state: () => ({}),
  actions: {
    async test() {
      const apiInstance = new MyApi(import.meta.env.VITE_BACKEND_URL);

      try {
        // Call the dashboard method
        await apiInstance.dashboard();
        console.log('Dashboard data retrieved successfully');
      } catch (error) {
        console.error('Error fetching dashboard data:', error);
      }
    }
  }
})
