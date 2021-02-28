param (
    [Parameter(Mandatory=$true, Position=0)]
    [ArgumentCompleter({
        param ( $commandName, $parameterName, $wordToComplete, $commandAst, $fakeBoundParameters )

        @("Application", "ApplicationGeneralisation") | where { $_ -like "$wordToComplete*" }
    })]
    [String] $Type,

    [Parameter(Mandatory=$true, Position=1)]
    [String] $Name,

    [Switch] $Expand
)

Import-Module ModellingAutomationLayer.SecondOrderFunctions

$AnanaModelDir = "C:\Temp\Patterns\Ikea Patterns\Ikea\seikdev3"

$Applications = $null
$ApplicationStacks = $null
$ApplicationGeneralisations = $null
$DependingApplications = $null
$DependingApplicationGeneralisations = $null
$ProcessingRealms = $null

pushd
switch ($Type) {
    "Application" {

        cd $AnanaModelDir
        $Applications = @( Get-Application $Name )

        if ($Expand) {
            $ApplicationNames = @( $Applications | select Name -Unique | foreach { $_.Name } )
            $AllApplications = @( Get-Application )
            $DependingApplications = @( $ApplicationNames | foreach { $Dependency = $_; $AllApplications | where Dependencies -contains $Dependency | where Name -notin $ApplicationNames } )

            cd ( Get-Item $AnanaModelDir ).Parent.FullName

            $ApplicationGeneralisationNames = @( $Applications | select ApplicationGeneralisation -Unique | foreach { $_.ApplicationGeneralisation } )
            $ApplicationGeneralisations = @( $ApplicationGeneralisationNames | Get-ApplicationGeneralisation )
        }

        break
    }

    "ApplicationGeneralisation" {

        cd ( Get-Item $AnanaModelDir ).Parent.FullName
        $ApplicationGeneralisations = @( Get-ApplicationGeneralisation $Name )

        if ($Expand) {
            $ApplicationGeneralisationNames = @( $ApplicationGeneralisations | select Name -Unique | foreach { $_.Name } )
            $AllApplicationGeneralisations = @( Get-ApplicationGeneralisation )
            $DependingApplicationGeneralisations = @( $ApplicationGeneralisationNames | foreach { $Dependency = $_; $AllApplicationGeneralisations | foreach { $_.Dependencies | where Application -eq $Dependency | where Application -notin $ApplicationGeneralisationNames } } )

            cd $AnanaModelDir
            $Applications = @( Get-Application $Name | where ApplicationGeneralisation -like $Name )
        }

        break
    }

    default {
        Write-Error "invalid value for parameter '-Type' - expected an Anana model object-type but found '$Type'"
    }
}
popd

$Result = @{
    Applications = $Applications
    DependingApplications = $DependingApplications
    ApplicationGeneralisations = $ApplicationGeneralisations
    DependingApplicationGeneralisations = $DependingApplicationGeneralisations
}

return $Result | ConvertTo-JSON
