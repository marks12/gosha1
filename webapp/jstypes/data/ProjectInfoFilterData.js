
	import {ProjectInfoFilter, ProjectInfoFilterFilter} from '../apiModel';

let projectInfoFilterData = function () {
    return {
        projectInfoFilterFilter: new ProjectInfoFilterFilter(),
        panel: {
            type: '',
            show: false,
            create: 'create',
            edit: 'edit'
        },
        panelHeaderCreate: 'Создать',
        panelHeaderEdit: 'Изменить',
        panelSubmitButtonTextCreate: 'Создать',
        panelSubmitButtonTextEdit: 'Изменить',
        currentProjectInfoFilterItem: {
            item: new ProjectInfoFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default projectInfoFilterData;

