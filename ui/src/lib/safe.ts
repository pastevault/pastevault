import { logger } from "$lib/server/logger";
import type { ServiceError } from '@grpc/grpc-js';

export type Safe<T> =
    | {
          error: false;
          data: T;
      }
    | {
          error: true;
          msg: string;
          fields?: {
              field: string;
              tag: string;
          }[];
      };

export declare function safe<T>(promise: Promise<T>): Promise<Safe<T>>;
export declare function safe<T>(fn: () => T): Safe<T>;

export function grpcSafe<T>(
	res: (value: Safe<T>) => void,
): (err: ServiceError | null, data: T | undefined) => void;


export function grpcSafe<T>(resolve: (res: Safe<T>) => void) {
	return (err: ServiceError | null, data: never) => {
		if (err) {
			logger.error(err);
			if (err.code === 3) {
				let fields = [];
				try {
					fields = JSON.parse(err.details);
				} catch (e) {
					return resolve({
						error: true,
						msg: err?.message || "Something went wrong",
					});
				}

				return resolve({
					error: true,
					msg: "Invalid argument",
					fields: fields,
				});
			}
			return resolve({
				error: true,
				msg: err?.message || "Something went wrong",
			});
		}
		if (!data) {
            logger.error("No data returned");
            return resolve({
                error: true,
                msg: "No data returned",
            });
        }
        resolve({ data, error: false });
	};
}
