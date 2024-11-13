import {
  DefaultApi,
  Configuration as PrivateConf,
  PlayersApi as PrivatePlayersApi,
  UsersApi as PrivateUsersApi
} from '@/repositories/clients/private';
import {
  Configuration as PublicConf,
  UsersApi as PublicUsersApi,
  type LoginRequest,
  type LoginResponseSchema
} from '@/repositories/clients/public';


import * as runtime from '@/repositories/clients/private';
import { resetStores } from './store';

export class MyApi extends runtime.BaseAPI {

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

  getPrivateConf() {
    const token = sessionStorage.getItem('session_token');

    return new PrivateConf({
      basePath: `${this.backendUrl}/api`,
      credentials: "include",
      accessToken: token ?? "",
    });
  }

  async refreshToken(): Promise<void> {

    const api = new PublicUsersApi(this.getPublicConf())

    try {
      const response = await api.refreshToken()
      sessionStorage.setItem('session_token', response.sessionToken);

    } catch (error) {
      throw new Error('Failed to refresh token');
    }
  }

  private async authenticatedRequest(requestFunction: () => Promise<runtime.ApiResponse<any>>): Promise<runtime.ApiResponse<any>> {
    try {
      return await requestFunction();
    } catch (error) {
      if (error instanceof runtime.ResponseError && error.response.status === 401) {
        try {
          // Refresh the token if the request is unauthorized
          await this.refreshToken();

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

  private async requestDashboardApi(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
    const api = new DefaultApi(this.getPrivateConf())

    try {
      await api.dashboardRaw(initOverrides)
    } catch (error) {
      if (error instanceof runtime.ResponseError) {
        throw new runtime.ResponseError(error.response, error.message)
      }
      throw new Error('Failed to dashboard');
    }

    // Assuming the response should be wrapped;
    return new runtime.VoidApiResponse(new Response);
  }

  /**
   * Public method to access the dashboard endpoint
   */
  async dashboard(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
    await this.authenticatedRequest(() => this.requestDashboardApi(initOverrides));
  }

  private async requestLoginApi(requestParameters: LoginRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<LoginResponseSchema> {
    const api = new PublicUsersApi(this.getPublicConf())

    try {
      return api.login(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof runtime.ResponseError) {
        throw new runtime.ResponseError(error.response, error.message)
      }
      throw new Error('Failed to login');
    }
  }

  /**
   * Public method to access the dashboard endpoint
   */
  async login(requestParameters: LoginRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<LoginResponseSchema> {
    return this.requestLoginApi(requestParameters, initOverrides);
  }

  private async requestLogoutApi(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
    const api = new PrivateUsersApi(this.getPrivateConf())

    try {
      await api.logout(initOverrides)
    } catch (error) {
      if (error instanceof runtime.ResponseError) {
        throw new runtime.ResponseError(error.response, error.message)
      }
      throw new Error('Failed to logout');
    }

    // Assuming the response should be wrapped;
    return new runtime.VoidApiResponse(new Response);
  }

  /**
   * Public method to access the dashboard endpoint
   */
  async logoutRequest(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
    await this.authenticatedRequest(() => this.requestLogoutApi(initOverrides));
  }

  /**
   * Logs out the user and redirects to the sign-in page
   */
  deleteSession(): void {
    // Clear session storage or any other storage that holds your authentication tokens
    sessionStorage.removeItem('session_token');

    resetStores()

    // Redirect the user to the sign-in page
    window.location.href = this.signInPageUrl;
  }

  private async requestCurrentUser(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<runtime.GetLoggedInUser200Response>> {
    const api = new PrivateUsersApi(this.getPrivateConf())

    try {
      return api.getLoggedInUserRaw(initOverrides)
    } catch (error) {
      if (error instanceof runtime.ResponseError) {
        throw new runtime.ResponseError(error.response, error.message)
      }
      throw new Error('Failed to get current user');
    }
  }

  /**
   * Public method to access the dashboard endpoint
   */
  async currentUser(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.GetLoggedInUser200Response> {
    const apiResponse = await this.authenticatedRequest(() => this.requestCurrentUser(initOverrides));
    return apiResponse.value();
  }

  private async listPlayersRequest(requestParameters: runtime.ListPlayersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<runtime.ListPlayers200Response>> {
    const api = new PrivatePlayersApi(this.getPrivateConf())

    try {
      return api.listPlayersRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof runtime.ResponseError) {
        throw new runtime.ResponseError(error.response, error.message)
      }
      throw new Error('Failed to list players');
    }
  }

  /**
   * Public method to access the dashboard endpoint
   */
  async listPlayers(requestParameters: runtime.ListPlayersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ListPlayers200Response> {
    const apiResponse = await this.authenticatedRequest(() => this.listPlayersRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async addPlayerRequest(requestParameters: runtime.AddPlayerRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<runtime.Player>> {
    const api = new PrivatePlayersApi(this.getPrivateConf())

    try {
      return api.addPlayerRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof runtime.ResponseError) {
        throw new runtime.ResponseError(error.response, error.message)
      }
      throw new Error('Failed to add player');
    }
  }

  async addPlayer(requestParameters: runtime.AddPlayerRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.Player> {
    const apiResponse = await this.authenticatedRequest(() => this.addPlayerRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async updatePlayerRequest(requestParameters: runtime.UpdatePlayerWithIdOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<runtime.Player>> {
    const api = new PrivatePlayersApi(this.getPrivateConf())

    try {
      return api.updatePlayerWithIdRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof runtime.ResponseError) {
        throw new runtime.ResponseError(error.response, error.message)
      }
      throw new Error('Failed to update player');
    }
  }

  async updatePlayer(requestParameters: runtime.UpdatePlayerWithIdOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.Player> {
    const apiResponse = await this.authenticatedRequest(() => this.updatePlayerRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async getPlayerRequest(requestParameters: runtime.GetPlayerWithIdRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<runtime.Player>> {
    const api = new PrivatePlayersApi(this.getPrivateConf())

    try {
      return api.getPlayerWithIdRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof runtime.ResponseError) {
        throw new runtime.ResponseError(error.response, error.message)
      }
      throw new Error('Failed to get player');
    }
  }

  async getPlayer(requestParameters: runtime.GetPlayerWithIdRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.Player> {
    const apiResponse = await this.authenticatedRequest(() => this.getPlayerRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async startGameRequest(requestParameters: runtime.StartGameRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<runtime.StartGame201Response>> {
    const api = new runtime.GameApi(this.getPrivateConf())

    try {
      return api.startGameRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof runtime.ResponseError) {
        throw new runtime.ResponseError(error.response, error.message)
      }
      throw new Error('Failed to start game');
    }
  }

  async startGame(requestParameters: runtime.StartGameRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.StartGame201Response> {
    const apiResponse = await this.authenticatedRequest(() => this.startGameRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async addGameStepsRequest(requestParameters: runtime.AddGameStepsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<runtime.AddGameSteps201Response>> {
    const api = new runtime.GameApi(this.getPrivateConf())

    try {
      return api.addGameStepsRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof runtime.ResponseError) {
        throw new runtime.ResponseError(error.response, error.message)
      }
      throw new Error('Failed to add game steps');
    }
  }

  async addGameSteps(requestParameters: runtime.AddGameStepsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.AddGameSteps201Response> {
    const apiResponse = await this.authenticatedRequest(() => this.addGameStepsRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async deleteGameStepsRequest(requestParameters: runtime.DeleteGameStepsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
    const api = new runtime.GameApi(this.getPrivateConf())

    try {
      return api.deleteGameStepsRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof runtime.ResponseError) {
        throw new runtime.ResponseError(error.response, error.message)
      }
      throw new Error('Failed to delete game steps');
    }
  }

  async deleteGameSteps(requestParameters: runtime.DeleteGameStepsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
    const apiResponse = await this.authenticatedRequest(() => this.deleteGameStepsRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }

  private async endGameRequest(requestParameters: runtime.EndGameOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
    const api = new runtime.GameApi(this.getPrivateConf())

    try {
      return api.endGameRaw(requestParameters, initOverrides)
    } catch (error) {
      if (error instanceof runtime.ResponseError) {
        throw new runtime.ResponseError(error.response, error.message)
      }
      throw new Error('Failed to end game');
    }
  }

  async endGame(requestParameters: runtime.EndGameOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
    const apiResponse = await this.authenticatedRequest(() => this.endGameRequest(requestParameters, initOverrides));
    return await apiResponse.value();
  }
}