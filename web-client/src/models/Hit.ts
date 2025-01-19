import type { Email } from "./Email";

export interface Hit {
    _index: string;
    _id: string;
    _source: {
        email: Email;
    };
}
