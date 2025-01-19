import type { Hit } from "./Hit";

export interface EmailResponse {
    hits: {
        hits: Hit[]
    };
}
