import type { EmailAddress } from "./Address.ts";

export interface Email {
    "Message-Id": string;
    "Date": string;
    "From": string;
    "To": EmailAddress[];
    "Cc": EmailAddress[];
    "Bcc": EmailAddress[];
    "Subject": string;
    "Mime-Version": string;
    "Content-Type": string;
    "content-Transfer-Encoding": string;
    "X-From": string;
    "X-To": EmailAddress[];
    "X-Cc": EmailAddress[];
    "X-Bcc": EmailAddress[];
    "X-Folder": string;
    "X-Origin": string;
    "X-FileName": string;
    "Body": string;
}
