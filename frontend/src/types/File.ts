import moment from "moment";
import { DATES_FORMAT } from ".";

export type Files = File[];
export class File {
    
    public readonly fileDate: moment.Moment;

    constructor(
        public readonly fileId: string,
        public readonly exists: string,
        fileDate: string
    ) {
        this.fileDate = moment(fileDate,DATES_FORMAT);
    }
}
  