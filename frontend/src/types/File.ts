import moment, { Moment } from "moment";
import { DATES_FORMAT, DATES_FORMAT_DISPLAY } from ".";

export type Files = File[];
export class File {

    private _fileDate: Moment;
    private _fileUrl: string | undefined;

    private constructor(
        public readonly fileId: string,
        public readonly exists: string,
        fileDate: string,
    ) {
        this._fileDate = moment(fileDate,DATES_FORMAT);
    }

    set fileDate(value: Moment | string | undefined) {
        if (typeof this._fileDate === 'string') {
            this._fileDate = moment(value, DATES_FORMAT);
        }  
        else if (value instanceof moment) {
            this._fileDate = value as moment.Moment;
        }
    }

    get fileDate(): Moment | undefined {
        if (typeof this._fileDate === 'string') {
            return moment(this._fileDate, DATES_FORMAT);
        }

        return this._fileDate;
    }

    getHumanReadableDate(): string  {
        return this.fileDate?.format(DATES_FORMAT_DISPLAY) || "";
    }

    set fileUrl(url: string | undefined) {
        this._fileUrl = url;
    }

    get fileUrl(): string | undefined {
        return this._fileUrl;
    }

    public static fromJson(json: any): File {
        return new File(
            json.fileId,
            json.exists,
            json.fileDate
        );
    }
}
