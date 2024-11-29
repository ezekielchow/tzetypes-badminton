import {
  BaseAPI,
  DefaultApi,
  GameApi,
  Configuration as PrivateConf,
  PlayersApi as PrivatePlayersApi,
  UsersApi as PrivateUsersApi,
  ResponseError,
  VoidApiResponse,
  type AddGameSteps201Response,
  type AddGameStepsRequest,
  type AddPlayerRequest,
  type ApiResponse,
  type CreateOrUpdateGameHistoryRequest,
  type DeleteGameStepsRequest,
  type EndGameOperationRequest,
  type GetGameHistory200Response,
  type GetGameHistoryRequest,
  type GetLoggedInUser200Response,
  type GetPlayerWithIdRequest,
  type GetRecentStatistics200Response,
  type InitOverrideFunction,
  type ListPlayers200Response,
  type ListPlayersRequest,
  type Player,
  type StartGame201Response,
  type StartGameRequest,
  type UpdatePlayerWithIdOperationRequest,
} from '@/repositories/clients/private';

import { MyPublicApi } from './requests-public';
import { resetStores } from './store';

export class MyPrivateApi extends BaseAPI {

  private signInPageUrl: string = '/login';
  private backendUrl: string;

  constructor(backendUrl: string) {
    super(); // Call to the base class constructor if needed.
    this.backendUrl = backendUrl;
  }

  getPrivateConf() {
    const token = sessionStorage.getItem('session_token');

    return new PrivateConf({
      basePath: `${this.backendUrl}/api`,
      credentials: "include",
      accessToken: token ?? "",
    });
  }

  private async authenticatedRequest(requestFunction: () => Promise<ApiResponse<any>>): Promise<ApiResponse<any>> {
    try {
      return await requestFunction();
    } catch (error) {
      if (error instanceof ResponseError && error.response.status === 401) {
        try {
          // Refresh the token if the request is unauthorized
          const publicApi = new MyPublicApi(this.backendUrl)
          await publicApi.refreshToken()

          // Retry the request after refreshing the token
          return await requestFunction();
        } catch (refreshError) {
          console.error('Token refresh failed: ', refreshError);

          // Logout and redirect to the sign-in page on failure
          this.deleteSession();
          throw new Error('Session expired. Please log in again.'); // Optionally rethrow to indicate the session issue
        }
      } else {
        // Rethrow the error if it's not a 401 or a refresh failure
        throw error;
      }
    }
  }

  private async requestDashboardApi(initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<void>> {
    const api = new DefaultApi(this.getPrivateConf())

    try {
      await api.dashboardRaw(initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to dashboard');
    }

    // Assuming the response should be wrapped;
    return new VoidApiResponse(new Response);
  }

  /**
   * Public method to access the dashboard endpoint
   */
  async dashboard(initOverrides?: RequestInit | InitOverrideFunction): Promise<void> {
    await this.authenticatedRequest(() => this.requestDashboardApi(initOverrides));
  }


  private async requestLogoutApi(initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<void>> {
    const api = new PrivateUsersApi(this.getPrivateConf())

    try {
      await api.logout(initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to logout');
    }

    // Assuming the response should be wrapped;
    return new VoidApiResponse(new Response);
  }

  /**
   * Public method to access the dashboard endpoint
   */
  async logoutRequest(initOverrides?: RequestInit | InitOverrideFunction): Promise<void> {
    await this.authenticatedRequest(() => this.requestLogoutApi(initOverrides));
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

  private async requestCurrentUser(initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<GetLoggedInUser200Response>> {
    const api = new PrivateUsersApi(this.getPrivateConf())

    try {
      return await api.getLoggedInUserRaw(initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to get current user');
    }
  }

  /**
   * Public method to access the dashboard endpoint
   */
  async currentUser(initOverrides?: RequestInit | InitOverrideFunction): Promise<GetLoggedInUser200Response> {
    const apiResponse = await this.authenticatedRequest(() => this.requestCurrentUser(initOverrides));
    return await apiResponse.value();
  }

  private async listPlayersRequest(requestParameters: ListPlayersRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<ListPlayers200Response>> {
    const api = new PrivatePlayersApi(this.getPrivateConf())

    try {
      return await api.listPlayersRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to list players');
    }
  }

  /**
   * Public method to access the dashboard endpoint
   */
  async listPlayers(requestParameters: ListPlayersRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ListPlayers200Response> {
    const apiResponse = await this.authenticatedRequest(() => this.listPlayersRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async addPlayerRequest(requestParameters: AddPlayerRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<Player>> {
    const api = new PrivatePlayersApi(this.getPrivateConf())

    try {
      return await api.addPlayerRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to add player');
    }
  }

  async addPlayer(requestParameters: AddPlayerRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<Player> {
    const apiResponse = await this.authenticatedRequest(() => this.addPlayerRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async updatePlayerRequest(requestParameters: UpdatePlayerWithIdOperationRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<Player>> {
    const api = new PrivatePlayersApi(this.getPrivateConf())

    try {
      return await api.updatePlayerWithIdRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to update player');
    }
  }

  async updatePlayer(requestParameters: UpdatePlayerWithIdOperationRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<Player> {
    const apiResponse = await this.authenticatedRequest(() => this.updatePlayerRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async getPlayerRequest(requestParameters: GetPlayerWithIdRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<Player>> {
    const api = new PrivatePlayersApi(this.getPrivateConf())

    try {
      return await api.getPlayerWithIdRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to get player');
    }
  }

  async getPlayer(requestParameters: GetPlayerWithIdRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<Player> {
    const apiResponse = await this.authenticatedRequest(() => this.getPlayerRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async startGameRequest(requestParameters: StartGameRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<StartGame201Response>> {
    const api = new GameApi(this.getPrivateConf())

    try {
      return await api.startGameRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to start game');
    }
  }

  async startGame(requestParameters: StartGameRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<StartGame201Response> {
    const apiResponse = await this.authenticatedRequest(() => this.startGameRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async addGameStepsRequest(requestParameters: AddGameStepsRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<AddGameSteps201Response>> {
    const api = new GameApi(this.getPrivateConf())

    try {
      return await api.addGameStepsRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to add game steps');
    }
  }

  async addGameSteps(requestParameters: AddGameStepsRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<AddGameSteps201Response> {
    const apiResponse = await this.authenticatedRequest(() => this.addGameStepsRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async deleteGameStepsRequest(requestParameters: DeleteGameStepsRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<void>> {
    const api = new GameApi(this.getPrivateConf())

    try {
      return await api.deleteGameStepsRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to delete game steps');
    }
  }

  async deleteGameSteps(requestParameters: DeleteGameStepsRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<void> {
    const apiResponse = await this.authenticatedRequest(() => this.deleteGameStepsRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async endGameRequest(requestParameters: EndGameOperationRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<void>> {
    const api = new GameApi(this.getPrivateConf())

    try {
      return await api.endGameRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to end game');
    }
  }

  async endGame(requestParameters: EndGameOperationRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<void> {
    const apiResponse = await this.authenticatedRequest(() => this.endGameRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async createOrUpdateGameHistoryRequest(requestParameters: CreateOrUpdateGameHistoryRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<GetGameHistory200Response>> {
    const api = new GameApi(this.getPrivateConf())

    try {
      return await api.createOrUpdateGameHistoryRaw(requestParameters, initOverrides)
    } catch (error: any) {
      if (error.response && error.message) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to create or update game history');
    }
  }

  async createOrUpdateGameHistory(requestParameters: CreateOrUpdateGameHistoryRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<GetGameHistory200Response> {
    const apiResponse = await this.authenticatedRequest(() => this.createOrUpdateGameHistoryRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async getGameHistoryRequest(requestParameters: GetGameHistoryRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<GetGameHistory200Response>> {
    const api = new GameApi(this.getPrivateConf())

    try {
      return await api.getGameHistoryRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to get game history');
    }
  }

  async getGameHistory(requestParameters: GetGameHistoryRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<GetGameHistory200Response> {
    const apiResponse = await this.authenticatedRequest(() => this.getGameHistoryRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async getRecentStatisticsRequest(initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<GetRecentStatistics200Response>> {
    const api = new GameApi(this.getPrivateConf())

    try {
      return await api.getRecentStatisticsRaw(initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to get recent statistics');
    }
  }

  async getRecentStatistics(initOverrides?: RequestInit | InitOverrideFunction): Promise<GetRecentStatistics200Response> {
    try {
      const apiResponse = await this.authenticatedRequest(() => this.getRecentStatisticsRequest(initOverrides));
      return apiResponse.value(); // Return the response value
    } catch (error) {
      console.error('Failed to fetch recent statistics:', error); // This will catch and log any error
      throw error; // Re-throw the error to propagate it further if needed
    }
  }
}