import {
  BaseAPI,
  GameApi,
  Configuration as PrivateConf,
  PlayersApi as PrivatePlayersApi,
  UsersApi as PrivateUsersApi,
  ResponseError,
  type AddGameSteps201Response,
  type AddGameStepsRequest,
  type AddPlayerRequest,
  type ApiResponse,
  type CreateOrUpdateGameHistoryRequest,
  type DeleteGameStepsRequest,
  type EndGameOperationRequest,
  type GetGame200Response,
  type GetGameHistory200Response,
  type GetGameHistoryRequest,
  type GetGameRequest,
  type GetLoggedInUser200Response,
  type GetPlayerWithIdRequest,
  type GetRecentStatistics200Response,
  type InitOverrideFunction,
  type ListActiveGames200Response,
  type ListPlayers200Response,
  type ListPlayersRequest,
  type Player,
  type StartGame201Response,
  type StartGameRequest,
  type UpdatePlayerWithIdOperationRequest
} from '@/repositories/clients/private';

import { useUserStore } from '@/stores/user-store';
import { resetStores } from './store';

export class MyPrivateApi extends BaseAPI {

  private signInPageUrl: string = '/login';
  private backendUrl: string;

  constructor(backendUrl: string) {
    super(); // Call to the base class constructor if needed.
    this.backendUrl = backendUrl;
  }

  decodeJWT(token: string) {
    const payload = token.split(".")[1]; // Extract the payload part
    const decodedPayload = JSON.parse(atob(payload)); // Base64 decode and parse JSON
    return decodedPayload;
  };

  async getPrivateConf() {
    const userStore = useUserStore()

    return new PrivateConf({
      basePath: `${this.backendUrl}/api`,
      credentials: "include",
      accessToken: userStore.firebaseIdToken,
    });
  }

  private async authenticatedRequest(
    requestFunction: () => Promise<ApiResponse<any>>
  ): Promise<ApiResponse<any> | void> {
    try {
      return await requestFunction();
    } catch (error: any) {
      if (error instanceof ResponseError && error.response.status === 401) {
        try {
          const userStore = useUserStore()
          const user = userStore.firebaseUser; // Replace with your Firebase auth instance
          if (user) {
            const newToken = await user.getIdToken(true); // Force refresh

            userStore.firebaseIdToken = newToken

            // Retry the original request
            return await requestFunction();
          }
        } catch (refreshError) {
          console.error("Failed to refresh token:", refreshError);
          this.deleteSession(); // Clear session if refreshing fails
          return;
        }
      }

      // Rethrow the error if it's not a 401 or a refresh failure
      throw error;
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

  private async requestCurrentUser(initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<GetLoggedInUser200Response>> {
    const api = new PrivateUsersApi(await this.getPrivateConf())

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
  async currentUser(initOverrides?: RequestInit | InitOverrideFunction): Promise<GetLoggedInUser200Response | void> {
    const apiResponse = await this.authenticatedRequest(() => this.requestCurrentUser(initOverrides));
    if (apiResponse) {
      return await apiResponse.value();
    }
    return
  }

  private async listPlayersRequest(requestParameters: ListPlayersRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<ListPlayers200Response>> {
    const api = new PrivatePlayersApi(await this.getPrivateConf())

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
  async listPlayers(requestParameters: ListPlayersRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ListPlayers200Response | void> {
    const apiResponse = await this.authenticatedRequest(() => this.listPlayersRequest(requestParameters, initOverrides));
    if (apiResponse) {
      return await apiResponse.value();
    }
    return
  }

  private async addPlayerRequest(requestParameters: AddPlayerRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<Player>> {
    const api = new PrivatePlayersApi(await this.getPrivateConf())

    try {
      return await api.addPlayerRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to add player');
    }
  }

  async addPlayer(requestParameters: AddPlayerRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<Player | void> {
    const apiResponse = await this.authenticatedRequest(() => this.addPlayerRequest(requestParameters, initOverrides));
    if (apiResponse) {
      return await apiResponse.value();
    }
    return
  }

  private async updatePlayerRequest(requestParameters: UpdatePlayerWithIdOperationRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<Player>> {
    const api = new PrivatePlayersApi(await this.getPrivateConf())

    try {
      return await api.updatePlayerWithIdRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to update player');
    }
  }

  async updatePlayer(requestParameters: UpdatePlayerWithIdOperationRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<Player | void> {
    const apiResponse = await this.authenticatedRequest(() => this.updatePlayerRequest(requestParameters, initOverrides));
    if (apiResponse) {
      return await apiResponse.value();
    }
    return
  }

  private async getPlayerRequest(requestParameters: GetPlayerWithIdRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<Player>> {
    const api = new PrivatePlayersApi(await this.getPrivateConf())

    try {
      return await api.getPlayerWithIdRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to get player');
    }
  }

  async getPlayer(requestParameters: GetPlayerWithIdRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<Player | void> {
    const apiResponse = await this.authenticatedRequest(() => this.getPlayerRequest(requestParameters, initOverrides));
    if (apiResponse) {
      return await apiResponse.value();
    }
    return
  }

  private async startGameRequest(requestParameters: StartGameRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<StartGame201Response>> {
    const api = new GameApi(await this.getPrivateConf())

    try {
      return await api.startGameRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to start game');
    }
  }

  async startGame(requestParameters: StartGameRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<StartGame201Response | void> {
    const apiResponse = await this.authenticatedRequest(() => this.startGameRequest(requestParameters, initOverrides));
    if (apiResponse) {
      return await apiResponse.value();
    }
    return
  }

  private async addGameStepsRequest(requestParameters: AddGameStepsRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<AddGameSteps201Response>> {
    const api = new GameApi(await this.getPrivateConf())

    try {
      return await api.addGameStepsRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to add game steps');
    }
  }

  async addGameSteps(requestParameters: AddGameStepsRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<AddGameSteps201Response | void> {
    const apiResponse = await this.authenticatedRequest(() => this.addGameStepsRequest(requestParameters, initOverrides));
    if (apiResponse) {
      return await apiResponse.value();
    }
    return
  }

  private async deleteGameStepsRequest(requestParameters: DeleteGameStepsRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<void>> {
    const api = new GameApi(await this.getPrivateConf())

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
    if (apiResponse) {
      return await apiResponse.value();
    }
    return
  }

  private async endGameRequest(requestParameters: EndGameOperationRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<void>> {
    const api = new GameApi(await this.getPrivateConf())

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
    if (apiResponse) {
      return await apiResponse.value();
    }
    return
  }

  private async createOrUpdateGameHistoryRequest(requestParameters: CreateOrUpdateGameHistoryRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<GetGameHistory200Response>> {
    const api = new GameApi(await this.getPrivateConf())

    try {
      return await api.createOrUpdateGameHistoryRaw(requestParameters, initOverrides)
    } catch (error: any) {
      if (error.response && error.message) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to create or update game history');
    }
  }

  async createOrUpdateGameHistory(requestParameters: CreateOrUpdateGameHistoryRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<GetGameHistory200Response | void> {
    const apiResponse = await this.authenticatedRequest(() => this.createOrUpdateGameHistoryRequest(requestParameters, initOverrides));
    if (apiResponse) {
      return await apiResponse.value();
    }
    return
  }

  private async getGameHistoryRequest(requestParameters: GetGameHistoryRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<GetGameHistory200Response>> {
    const api = new GameApi(await this.getPrivateConf())

    try {
      return await api.getGameHistoryRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to get game history');
    }
  }

  async getGameHistory(requestParameters: GetGameHistoryRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<GetGameHistory200Response | void> {
    const apiResponse = await this.authenticatedRequest(() => this.getGameHistoryRequest(requestParameters, initOverrides));
    if (apiResponse) {
      return await apiResponse.value();
    }
    return
  }

  private async getRecentStatisticsRequest(initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<GetRecentStatistics200Response>> {
    const api = new GameApi(await this.getPrivateConf())

    try {
      return await api.getRecentStatisticsRaw(initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to get recent statistics');
    }
  }

  async getRecentStatistics(initOverrides?: RequestInit | InitOverrideFunction): Promise<GetRecentStatistics200Response | void> {
    try {
      const apiResponse = await this.authenticatedRequest(() => this.getRecentStatisticsRequest(initOverrides));
      if (apiResponse) {
        return await apiResponse.value();
      }
      return
    } catch (error) {
      console.error('Failed to fetch recent statistics:', error); // This will catch and log any error
      throw error; // Re-throw the error to propagate it further if needed
    }
  }

  private async listActiveGamesRequest(initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<ListActiveGames200Response>> {
    const api = new GameApi(await this.getPrivateConf())

    try {
      return await api.listActiveGamesRaw(initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to get active games');
    }
  }

  async listActiveGames(initOverrides?: RequestInit | InitOverrideFunction): Promise<ListActiveGames200Response | void> {
    try {
      const apiResponse = await this.authenticatedRequest(() => this.listActiveGamesRequest(initOverrides));
      if (apiResponse) {
        return await apiResponse.value();
      }
      return
    } catch (error) {
      console.error('Failed to fetch active games:', error); // This will catch and log any error
      throw error; // Re-throw the error to propagate it further if needed
    }
  }

  private async getGameRequest(requestParameters: GetGameRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<ApiResponse<GetGame200Response>> {
    const api = new GameApi(await this.getPrivateConf())

    try {
      return await api.getGameRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof ResponseError) {
        throw new ResponseError(error.response, error.message)
      }
      throw new Error('Failed to get game statistics');
    }
  }

  async getGame(requestParameters: GetGameRequest, initOverrides?: RequestInit | InitOverrideFunction): Promise<GetGame200Response | void> {
    const apiResponse = await this.authenticatedRequest(() => this.getGameRequest(requestParameters, initOverrides));
    if (apiResponse) {
      return await apiResponse.value();
    }
    return
  }
}