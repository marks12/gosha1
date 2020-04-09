
	import {ProjectInfo, ProjectInfoFilter} from '../apiModel';

let projectInfoData = function () {
    return {
        projectInfoFilter: new ProjectInfoFilter(),
        panel: {
            type: '',
            show: false,
            create: 'create',
            edit: 'edit',
            request: 'request'
        },
        panelHeaderCreate: 'Создать',
        panelHeaderEdit: 'Изменить',
        panelSubmitButtonTextCreate: 'Создать',
        panelSubmitButtonTextEdit: 'Изменить',
        currentProjectInfoItem: {
            item: new ProjectInfo(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default projectInfoData;

