import {
  DefaultApi as PrivateApi,
  Configuration as PrivateConf,
} from '@/repositories/clients/private';
import {
  DefaultApi as PublicApi,
  Configuration as PublicConf,
  type LoginRequest,
  type LoginResponseSchema
} from '@/repositories/clients/public';


import * as runtime from '@/repositories/clients/private';

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
      headers: { Authorization: `Bearer ${token}` },
      credentials: "include"
    });
  }

  async refreshToken(): Promise<void> {

    const api = new PublicApi(this.getPublicConf())

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
          this.logoutRequest()
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
    const api = new PrivateApi(this.getPrivateConf())

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
    const api = new PublicApi(this.getPublicConf())

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
    const api = new PrivateApi(this.getPrivateConf())

    try {
      await api.logoutRaw(initOverrides)
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

    // Redirect the user to the sign-in page
    window.location.href = this.signInPageUrl;
  }
}