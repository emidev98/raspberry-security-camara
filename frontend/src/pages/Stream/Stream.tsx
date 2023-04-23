import { useEffect, useRef, useState } from 'react'
import useApi from '../../hooks/useApi'
import { File } from '../../types';
import { BigPlayButton, ControlBar, LoadingSpinner, Player, PlayerReference, PlayerState } from 'video-react';
import "./Stream.scss"
import PauseIcon from '@mui/icons-material/Pause';
import PlayArrowIcon from '@mui/icons-material/PlayArrow';
import RefreshIcon from '@mui/icons-material/Refresh';

export const Stream : React.FC = () => {
    const api = useApi();
    const player = useRef<PlayerReference>(null);
    const [file, setFile] = useState<File>();
    const [lastSecond, setLastSecond] = useState<number>(0);
    const [playerState, setPlayerState] = useState<PlayerState>();

    useEffect(() => {
        const init = () => {
            api.downloadLatestFile().then(setFile);
            player.current?.subscribeToStateChange(setPlayerState)
        }
        init();
        const interval = setInterval(init, 1000 * 30);

        return () => clearInterval(interval);
    }, [])

    useEffect(()=>{
        if(file?.fileDate){
            let time = Number(playerState?.currentTime.toFixed());

            if(time !== lastSecond) {
                file.fileDate = file.fileDate.add(1, "seconds");
                setLastSecond(time);
            }
        }
    },[playerState]);

    const togglePause = () => {
        if(playerState?.paused) player.current?.play()
        else player.current?.pause()
    }

    const refresh = () => {
        player.current?.load()
        if (file?.fileDate) {
            file.fileDate = file?.fileDate.add(-lastSecond, "seconds");
            setLastSecond(0);
        }
    }

    return (
        <div id="StreamPage">
            <div className="Header">
                <h3>Stream</h3>
            </div>

            <Player src={file?.fileUrl} ref={player} autoPlay>
                <LoadingSpinner />
                <BigPlayButton position="center" />
                <ControlBar disableDefaultControls
                    disableCompletely/>
            </Player>

            <div className='Controls'>
                <div className="ControlsPause"
                    onClick={togglePause}>
                    {playerState?.paused 
                        ? <PlayArrowIcon fontSize='inherit'/>
                        : <PauseIcon fontSize='inherit'/>}
                </div>

                <div className='ControlsTime'>{file?.getHumanReadableDate()}</div>

                <div className="ControlsRefresh"
                    onClick={refresh}>
                    <RefreshIcon/>
                </div>
            </div>
        </div>
    )
}
