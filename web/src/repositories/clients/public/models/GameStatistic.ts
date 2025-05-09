/* tslint:disable */
/* eslint-disable */
/**
 * tzetypes-badminton
 * tzetypes-badminton
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface GameStatistic
 */
export interface GameStatistic {
    /**
     * 
     * @type {string}
     * @memberof GameStatistic
     */
    totalGameTime: string;
    /**
     * 
     * @type {string}
     * @memberof GameStatistic
     */
    leftConsecutivePoints: string;
    /**
     * 
     * @type {string}
     * @memberof GameStatistic
     */
    rightConsecutivePoints: string;
    /**
     * 
     * @type {string}
     * @memberof GameStatistic
     */
    leftLongestPoint: string;
    /**
     * 
     * @type {string}
     * @memberof GameStatistic
     */
    leftShortestPoint: string;
    /**
     * 
     * @type {string}
     * @memberof GameStatistic
     */
    rightLongestPoint: string;
    /**
     * 
     * @type {string}
     * @memberof GameStatistic
     */
    rightShortestPoint: string;
    /**
     * 
     * @type {string}
     * @memberof GameStatistic
     */
    averagePerPoint: string;
    /**
     * 
     * @type {string}
     * @memberof GameStatistic
     */
    leftAveragePerPoint: string;
    /**
     * 
     * @type {string}
     * @memberof GameStatistic
     */
    rightAveragePerPoint: string;
    /**
     * 
     * @type {string}
     * @memberof GameStatistic
     */
    consecutivePointsRatio: string;
    /**
     * 
     * @type {string}
     * @memberof GameStatistic
     */
    longestPointRatio: string;
    /**
     * 
     * @type {string}
     * @memberof GameStatistic
     */
    shortestPointRatio: string;
    /**
     * 
     * @type {string}
     * @memberof GameStatistic
     */
    averagePerPointRatio: string;
}

/**
 * Check if a given object implements the GameStatistic interface.
 */
export function instanceOfGameStatistic(value: object): value is GameStatistic {
    if (!('totalGameTime' in value) || value['totalGameTime'] === undefined) return false;
    if (!('leftConsecutivePoints' in value) || value['leftConsecutivePoints'] === undefined) return false;
    if (!('rightConsecutivePoints' in value) || value['rightConsecutivePoints'] === undefined) return false;
    if (!('leftLongestPoint' in value) || value['leftLongestPoint'] === undefined) return false;
    if (!('leftShortestPoint' in value) || value['leftShortestPoint'] === undefined) return false;
    if (!('rightLongestPoint' in value) || value['rightLongestPoint'] === undefined) return false;
    if (!('rightShortestPoint' in value) || value['rightShortestPoint'] === undefined) return false;
    if (!('averagePerPoint' in value) || value['averagePerPoint'] === undefined) return false;
    if (!('leftAveragePerPoint' in value) || value['leftAveragePerPoint'] === undefined) return false;
    if (!('rightAveragePerPoint' in value) || value['rightAveragePerPoint'] === undefined) return false;
    if (!('consecutivePointsRatio' in value) || value['consecutivePointsRatio'] === undefined) return false;
    if (!('longestPointRatio' in value) || value['longestPointRatio'] === undefined) return false;
    if (!('shortestPointRatio' in value) || value['shortestPointRatio'] === undefined) return false;
    if (!('averagePerPointRatio' in value) || value['averagePerPointRatio'] === undefined) return false;
    return true;
}

export function GameStatisticFromJSON(json: any): GameStatistic {
    return GameStatisticFromJSONTyped(json, false);
}

export function GameStatisticFromJSONTyped(json: any, ignoreDiscriminator: boolean): GameStatistic {
    if (json == null) {
        return json;
    }
    return {
        
        'totalGameTime': json['total_game_time'],
        'leftConsecutivePoints': json['left_consecutive_points'],
        'rightConsecutivePoints': json['right_consecutive_points'],
        'leftLongestPoint': json['left_longest_point'],
        'leftShortestPoint': json['left_shortest_point'],
        'rightLongestPoint': json['right_longest_point'],
        'rightShortestPoint': json['right_shortest_point'],
        'averagePerPoint': json['average_per_point'],
        'leftAveragePerPoint': json['left_average_per_point'],
        'rightAveragePerPoint': json['right_average_per_point'],
        'consecutivePointsRatio': json['consecutive_points_ratio'],
        'longestPointRatio': json['longest_point_ratio'],
        'shortestPointRatio': json['shortest_point_ratio'],
        'averagePerPointRatio': json['average_per_point_ratio'],
    };
}

export function GameStatisticToJSON(value?: GameStatistic | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'total_game_time': value['totalGameTime'],
        'left_consecutive_points': value['leftConsecutivePoints'],
        'right_consecutive_points': value['rightConsecutivePoints'],
        'left_longest_point': value['leftLongestPoint'],
        'left_shortest_point': value['leftShortestPoint'],
        'right_longest_point': value['rightLongestPoint'],
        'right_shortest_point': value['rightShortestPoint'],
        'average_per_point': value['averagePerPoint'],
        'left_average_per_point': value['leftAveragePerPoint'],
        'right_average_per_point': value['rightAveragePerPoint'],
        'consecutive_points_ratio': value['consecutivePointsRatio'],
        'longest_point_ratio': value['longestPointRatio'],
        'shortest_point_ratio': value['shortestPointRatio'],
        'average_per_point_ratio': value['averagePerPointRatio'],
    };
}

