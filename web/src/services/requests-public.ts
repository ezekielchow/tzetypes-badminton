import {
  BaseAPI,
  GameApi,
  Configuration as PublicConf,
  ResponseError,
  type ApiResponse,
  type GetGameStatistics200Response,
  type GetGameStatisticsRequest,
  type InitOverrideFunction
} from '@/repositories/clients/public';

import { resetStores } from './store';

export class MyPublicApi extends BaseAPI {

  private signInPageUrl: string = '/login';
  private backendUrl: string;

  constructor(backendUrl: string) {
    super(); // Call to the base class constructor if needed.
    this.backendUrl = backendUrl;
  }

  getPublicConf() {
    return new PublicConf({
      basePath: `${this.backendUrl}`,
      credentials: "include"
    });
  }

  private async authenticatedRequest(requestFunction: () => Promise<ApiResponse<any>>): Promise<ApiResponse<any>> {
    try {
      return await requestFunction();
    } catch (error) {
      if (error instanceof ResponseError && error.response.status === 401) {
        this.deleteSession();
        throw new Error('Session expired. Please log in again.'); // Optionally rethrow to indicate the session issue
      } else {
        // Rethrow the error if it's not a 401 or a refresh failure
        throw error;
      }
    }
  }

  /**
   * Logs out the user and redirects to the sign-in page
   */
  deleteSession(): void {
    // Clear session storage or any other storage that holds your authentication tokens
    sessionStorage.clear()
    localStorage.clear()

    resetStores()

    // Redirect the user to the sign-in page
    window.location.href = this.signInPageUrl;
  }

  private async getGameStatisticsRequest(requestParameters: GetGameStatisticsRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<GetGameStatistics200Response>> {
    const api = new GameApi(this.getPublicConf())

    try {
      return await api.getGameStatisticsRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to get game statistics');
    }
  }

  async getGameStatistics(requestParameters: GetGameStatisticsRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<GetGameStatistics200Response> {
    const apiResponse = await this.authenticatedRequest(() => this.getGameStatisticsRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }
}